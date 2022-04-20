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
	debugPtr := flag.Bool("debug", false, "Display any debugging information")
	servicePtr := flag.String("service", "", "Name service to be requested")
	filterPtr := flag.String("filter", "", "Filter parameter")
	filterValuePtr := flag.String("filter-value", "", "Filter value parameter")

	flag.Parse()

	// Flags Input validations, service is mandatory
	if *servicePtr == "" {
		log.Fatal("FATAL: service option is mandatory")
	}

	// Config
	viper.SetConfigName("client")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: server \n %v", err)
	}

	// Set up a connection to the server
	addr := fmt.Sprintf("%v:%d", viper.GetString("client.host_target"), viper.GetInt("client.port"))
	if *debugPtr {
		log.Printf("Connecting to %v", addr)
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}(conn)

	// Setup client
	c := client.NewClient(conn)

	// Run services
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
