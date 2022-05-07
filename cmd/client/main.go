package main

import (
	"flag"
	"fmt"
	config "github.com/marcos-wz/capstone/configs"
	"google.golang.org/grpc/credentials"
	"log"

	"github.com/marcos-wz/capstone/internal/client"
	"google.golang.org/grpc"
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
	clientConfig, err := config.LoadClientConfig("./configs/client.json")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}
	// SSL
	creds, sslErr := credentials.NewClientTLSFromFile(clientConfig.SSLCert, "")
	if sslErr != nil {
		log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
	}
	// Set up a connection to the server
	address := fmt.Sprintf("%v:%d", clientConfig.ServerHost, clientConfig.ServerPort)
	if *debugPtr {
		log.Printf("Connecting to %v", address)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
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
