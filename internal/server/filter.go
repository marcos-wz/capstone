package server

import (
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"io/fs"
	"log"
)

// Filter gets fruits filtered from the service. Only valid filters and values are allowed.
// PARAMS:
// 	- filter: the field filter request - filters allowed: id, name, color, country
//	- value: the filter value request
// RESPONSES:
// 	- 200 Status OK: returns filter response with fruits filtered list
//	- 422 Unprocessable Entity : returns param filter and value errors
//	- 500 Internal Server : returns reader CSV File error (critical!)
//	- 400 Bad Request: default errors
func (s *server) Filter(filterReq *filterpb.FilterRequest, stream fruitpb.FruitService_FilterServer) error {
	if Debug {
		log.Println("RPC Filter: starting...")
		log.Printf("RPC Filter: filter request: %+v", filterReq)
	}
	// SERVICE
	fruits, err := s.service.GetFilteredFruits(filterReq)
	// Error validations
	if err != nil {
		log.Printf("Error: %T - %v", err, err)

		switch err.(type) {
		case *fs.PathError:
			log.Printf("REPO ERROR: %v", err)
			return err
		default:
			return err
		}

		//switch err.Type {
		//// Repository File Error response: internal server error
		//case "Repo.FileError":
		//	return c.JSON(http.StatusInternalServerError, &entity.ErrorResponse{
		//		Message: err.Error.Error(),
		//	})
		//// Repository parser error response : partial fruits with parser errors
		//case "Repo.ParserError":
		//	return c.JSON(http.StatusPartialContent, &entity.FruitFilterResponse{
		//		Fruits:       fruits,
		//		ParserErrors: err.ParserErrors,
		//	})
		//default:
		//	// Default error response
		//	return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
		//		Message: err.Error.Error(),
		//	})
		//}
	}
	// Response
	log.Printf("Fruits found: %d \n%v", len(fruits), fruits)
	for _, fruit := range fruits {
		res := &filterpb.FilterResponse{Fruit: fruit}
		err := stream.Send(res)
		if err != nil {
			log.Printf("ERROR: sending stream: %v", err)
		}
	}
	return nil
}
