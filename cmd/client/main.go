package main

import (
	"flag"
	"fmt"
	"log"

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

	// Flags Input validations: service is mandatory
	if *servicePtr == "" {
		log.Fatal("FATAL: service option is mandatory")
	}

	// Config
	clientConfig, err := client.LoadClientConfig("./config/client.json")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	// Set up a connection to the server
	address := fmt.Sprintf("%v:%d", clientConfig.ServerHost, clientConfig.ServerPort)
	if *debugPtr {
		log.Printf("Connecting to %v", address)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	c := client.NewFruitClient(conn)

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
