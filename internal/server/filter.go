package server

import (
	"fmt"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"log"
)

// Filter streams filtered fruits from the filter service. Only valid filters and values are allowed.
func (s *server) Filter(filterReq *filterpb.FilterRequest, stream fruitpb.FruitService_FilterServer) error {
	if Debug {
		log.Println("RPC Filter: starting...")
		log.Printf("RPC Filter: filter request: %+v", filterReq)
	}
	// SERVICE
	fruits, err := s.service.FilterFruits(filterReq)
	// Error responses
	if err != nil {
		return err
	}
	// Response
	log.Printf("Filtered fruits found: %d", len(fruits))
	for _, fruit := range fruits {
		log.Printf("Sending fruit: %v ...", fruit)
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
