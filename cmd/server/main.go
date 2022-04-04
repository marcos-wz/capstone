package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/marcos-wz/capstone/internal/fruitpb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedFruitServiceServer
}

func (*server) Filter(ctx context.Context, req *pb.FilterRequest) (*pb.FilterResponse, error) {
	log.Printf("Filter Request: %v", req)
	resp := &pb.FilterResponse{
		Code: 200,
		Fruits: []*pb.Fruit{
			{
				Id:          1,
				Name:        "Pera",
				Description: "Fruta tropical",
				Color:       "green",
				// Unit: "lb",0,0,0,Canada,2022-02-01T12:14:05-06:00
			},
		},
	}
	return resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFruitServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
