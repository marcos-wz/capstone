package main

import (
	"fmt"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/marcos-wz/capstone/internal/server"
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
	//readerRepo := repository.NewReaderRepo(viper.GetString("data.csv"))
	//filterSvc := service.NewFilterService(readerRepo)
	//
	//serverServices := server.NewServer(filterSvc, nil, nil)
	srv := server.NewServer()

	// Listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("server.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Server
	s := grpc.NewServer()
	fruitpb.RegisterFruitServiceServer(s, srv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
