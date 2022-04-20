package server

import (
	"context"
	"github.com/marcos-wz/capstone/proto/loaderpb"
	"log"
)

func (*server) Loader(ctx context.Context, in *loaderpb.LoaderRequest) (*loaderpb.LoaderResponse, error) {
	log.Println("RPC Loader starting...")
	return nil, nil
}
