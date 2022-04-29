package main

import (
	"fmt"
	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/internal/service"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/marcos-wz/capstone/internal/server"
)

func main() {

	// Load Config
	serverConfig, err := server.LoadServerConfig("./config/server.json")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}
	// SSL
	creds, sslErr := credentials.NewServerTLSFromFile(serverConfig.SSLCert, serverConfig.SSLKey)
	if sslErr != nil {
		log.Fatalf("Failed loading certificates: %v", sslErr)
	}
	// Fruit repository
	repo := repository.NewFruitRepo(serverConfig.CSVFile, serverConfig.ExternalAPI)
	// Fruit service
	svc := service.NewFruitService(repo)
	// Fruit server
	fruitServer := server.NewFruitServer(svc)

	// Listener
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", serverConfig.Host, serverConfig.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// GRPC server
	s := grpc.NewServer(grpc.Creds(creds))
	fruitpb.RegisterFruitServiceServer(s, fruitServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
