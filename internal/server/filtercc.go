package server

import (
	"github.com/marcos-wz/capstone/proto/filterccpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"log"
)

func (*server) FilterCC(req *filterccpb.FilterCCRequest, stream fruitpb.FruitService_FilterCCServer) error {
	log.Println("RPC Filter CC starting...")
	return nil
}
