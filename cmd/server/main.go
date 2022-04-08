package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/marcos-wz/capstone/internal/fruitpb"
	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/internal/server"
	"github.com/marcos-wz/capstone/internal/service"
	"github.com/spf13/viper"
)

func main() {

	// Config
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: server \n %v", err)
	}

	// Server Services
	readerRepo := repository.NewReaderRepo(viper.GetString("data.csv"))
	filterSvc := service.NewFilterService(readerRepo)

	serverServices := server.NewServerServices(filterSvc, nil, nil)

	// Listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("server.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Server
	s := grpc.NewServer()
	pb.RegisterFruitServiceServer(s, serverServices)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
