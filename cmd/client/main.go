package main

import (
	"flag"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/marcos-wz/capstone/internal/fruitpb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFruitServiceClient(conn)

	//Conact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Filter(ctx, &pb.FilterRequest{
		Filter: pb.FiltersAllowed_COLOR,
		Value:  "green",
	})
	if err != nil {
		log.Fatalf("could not filter: %v", err)
	}
	log.Printf("Resp Code: %d", r.GetCode())
	log.Printf("Resp Fruits: %s", r.GetFruits())
	// log.Printf("Resp Code: %d", r.GetCode())
}
