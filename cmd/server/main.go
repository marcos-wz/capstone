package main

import (
	"fmt"
	config "github.com/marcos-wz/capstone/configs"

	"google.golang.org/grpc/credentials"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/internal/service"
	"github.com/marcos-wz/capstone/proto/fruitpb"
)

func main() {

	// Load Config
	serverConfig, err := config.LoadServerConfig("./configs/server.json")
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
	// Fruit services
	services := service.NewFruitService(repo)

	// Listener
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", serverConfig.Host, serverConfig.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// GRPC server
	s := grpc.NewServer(grpc.Creds(creds))
	fruitpb.RegisterFruitServiceServer(s, services)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
