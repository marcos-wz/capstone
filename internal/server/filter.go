package server

import (
	"fmt"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/fs"
	"log"
)

// Filter streams filtered fruits from the filter service. Only valid filters and values are allowed.
// PARAMS:
// 	- filter: the filter request - filters allowed: id, name, color, country
//	- value: the filter value request
// RESPONSES:
// 	- returns filter response with filtered fruit
//	- Data Loss: returns reader CSV, and fetcher JSON Files error (critical!)
//	- Unknown: default errors
//	- Internal: sending stream errors
func (s *server) Filter(filterReq *filterpb.FilterRequest, stream fruitpb.FruitService_FilterServer) error {
	if Debug {
		log.Println("RPC Filter: starting...")
		log.Printf("RPC Filter: filter request: %+v", filterReq)
	}
	// SERVICE
	fruits, err := s.service.GetFilteredFruits(filterReq)
	// Error responses
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			return status.Errorf(
				codes.DataLoss,
				fmt.Sprintf("Repository error file: %v", err),
			)
		default:
			return status.Errorf(
				codes.Unknown,
				fmt.Sprintf("%v", err),
			)
		}
	}
	// Response
	log.Printf("Fruits found: %d \n%v", len(fruits), fruits)
	for _, fruit := range fruits {
		res := &filterpb.FilterResponse{Fruit: fruit}
		err := stream.Send(res)
		if err != nil {
			errStream := fmt.Errorf("sending stream: %v", err)
			log.Printf("ERROR: %v", errStream)
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("%v", errStream),
			)
		}
	}
	return nil
}
