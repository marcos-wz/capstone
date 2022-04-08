package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/marcos-wz/capstone/internal/client"
)

func main() {
	// Command Line Flags
	servicePtr := flag.String("service", "", "the service to be requested")
	filterPtr := flag.String("filter", "", "the filter parameter")
	filterValuePtr := flag.String("filter-value", "", "the filter value parameter")
	flag.Parse()

	// Config
	viper.SetConfigName("client")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: server \n %v", err)
	}

	// Set up a connection to the server
	addr := fmt.Sprintf("%v:%d", viper.GetString("client.host_target"), viper.GetInt("client.port"))
	log.Printf("Connecting to %v", addr)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Setup client
	c := client.NewClient(conn)

	// Services
	switch *servicePtr {
	case "filter":
		c.Filter(*filterPtr, *filterValuePtr)
	case "loader":
		c.Loader()
	case "filterCC":
		c.FilterCC(*filterPtr, *filterValuePtr)
	default:
		log.Fatalf("FATAL: service %q not found", *servicePtr)
	}
}
