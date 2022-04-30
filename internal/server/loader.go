package server

import (
	"context"
	"github.com/marcos-wz/capstone/proto/loaderpb"
	"log"
)

func (*server) Loader(ctx context.Context, req *loaderpb.LoaderRequest) (*loaderpb.LoaderResponse, error) {
	log.Println("RPC Loader starting...")
	log.Printf("RPC Loader request: %v", req)
	log.Printf("RPC Loader context: %v", ctx)
	return nil, nil
}
