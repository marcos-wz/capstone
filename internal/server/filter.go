package server

import (
	"context"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"log"
)

// Filter gets fruits filtered from the service. Only valid filters and values are allowed.
// PARAMS:
// 	- filter: the field filter request - filters allowed: id, name, color, country
//	- value: the filter value request
// RESPONSES:
// 	- 200 Status OK: returns filter response with fruits filtered list
//	- 206 Partial Content: returns fruits filtered  partial list and reader parser errors(invalid csv file data!!)
//	- 422 Unprocessable Entity : returns param filter and value errors
//	- 500 Internal Server : returns reader CSV File error (critical!)
//	- 400 Bad Request: default errors
func (s *server) Filter(ctx context.Context, filterReq *filterpb.FilterRequest) (*filterpb.FilterResponse, error) {
	log.Println("RPC Filter: starting...")
	log.Printf("RPC Filter: filter request: %+v", filterReq)
	s.service.GetFilteredFruits(filterReq)
	// *************************************
	return &filterpb.FilterResponse{
		Fruit: &basepb.Fruit{
			Name: "fruit test",
		},
	}, nil
}

//func (s *server) Filter(ctx context.Context, req *pb.FilterRequest) (*pb.FilterResponse, error) {

//	// Filter request validation. Invalid filter response unprocessable entity
//	filter := &entity.FruitsFilterParams{
//		Filter: req.GetFilter().String(),
//		Value:  req.GetValue(),
//	}
//
//	if err := validator.New().Struct(filter); err != nil {
//		log.Println("ERROR: filter request validation - ", err)
//		return &pb.FilterResponse{
//			Code:  422,
//			Error: err.Error(),
//		}, fmt.Errorf("unprocessable entity: %v", err)
//	}
//
//	// Get Fruits from service
//	fruits, err := s.filterService.GetFilteredFruits(filter)
//	// Error validations
//	if err != nil {
//		switch err.Type {
//		// Repository File Error response: internal server error
//		case "Repo.FileError":
//			return &pb.FilterResponse{
//				Code:  500,
//				Error: err.Error.Error(),
//			}, fmt.Errorf("repository file error: %v", err.Error)
//		// Repository parser error response : partial fruits with parser errors
//		case "Repo.ParserError":
//			// Parser records
//			// for _, e := range err.ParserErrors {
//			// 	log.Printlnf(.Record)
//			// }
//			return &pb.FilterResponse{
//				Code:  206,
//				Error: err.Error.Error(),
//			}, fmt.Errorf("repository parser error: %v", err.Error)
//		default:
//			return &pb.FilterResponse{
//				Code:  404,
//				Error: err.Error.Error(),
//			}, fmt.Errorf("bad request: %v", err.Error)
//		}
//	}
//
//	// Mapping fruits from to Protocol Buffers
//	// log.Println("Fruits: ", fruits)
//	fruitsPb := []*pb.Fruit{}
//	for _, f := range fruits {
//		fruitsPb = append(fruitsPb, &pb.Fruit{
//			Id:          uint32(f.ID),
//			Name:        f.Name,
//			Description: f.Description,
//			Color:       f.Color,
//			Price:       float32(f.Price),
//			Stock:       uint32(f.Stock),
//			Country:     f.Country,
//		})
//	}
//	resp := &pb.FilterResponse{
//		Code:   200,
//		Fruits: fruitsPb,
//	}
//	return resp, nil
//}
