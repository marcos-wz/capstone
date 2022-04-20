package server

import (
	"context"
	"github.com/marcos-wz/capstone/proto/filterccpb"
	"log"
)

func (*server) FilterCC(ctx context.Context, in *filterccpb.FilterCCRequest) (*filterccpb.FilterCCResponse, error) {
	log.Println("RPC Filter CC starting...")
	return nil, nil
}
