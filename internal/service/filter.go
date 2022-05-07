package service

import (
	"fmt"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"log"
)

func (fs *fruitService) Filter(req *filterpb.FilterRequest, stream fruitpb.FruitService_FilterServer) error {
	if DebugLevel >= 1 {
		log.Println("RPC Filter starting...")
		log.Printf("RPC Filter request: %v", req)
	}
	// Repository
	fruits, err := fs.repo.ReadFruits()
	if err != nil {
		return err
	}
	// Filter Fruit List
	fruits, err = filterFactory(fruits, req)
	if err != nil {
		return err
	}
	// Response
	if DebugLevel >= 1 {
		log.Printf("Filtered fruits found: %d", len(fruits))
	}
	for _, fruit := range fruits {
		if DebugLevel >= 1 {
			log.Printf("Sending fruit: %v ...", fruit)
		}
		res := &filterpb.FilterResponse{Fruit: fruit}
		err := stream.Send(res)
		if err != nil {
			errStream := fmt.Errorf("sending stream: %v", err)
			log.Printf("ERROR: %v", errStream)
			return err
		}
	}
	return nil
}
