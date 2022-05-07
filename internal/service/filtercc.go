package service

import (
	"github.com/marcos-wz/capstone/proto/filterccpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"log"
)

func (*fruitService) FilterCC(req *filterccpb.FilterCCRequest, stream fruitpb.FruitService_FilterCCServer) error {
	log.Println("RPC Filter CC starting...")
	log.Printf("RPC Filter CC request: %v", req)
	log.Printf("RPC Filter CC stream: %v", stream)
	return nil
}
