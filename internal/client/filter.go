package client

import (
	"context"
	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"io"
	"log"
)

func (fc *fruitClient) Filter(filter, value string) {
	log.Println("Server Streaming RPC: starting...")
	log.Printf("Requesting Filter: %q, Value: %q ...", filter, value)
	// Input Validation
	filterPB := repository.ParseFilter(filter)
	if filterPB == filterpb.FiltersAllowed_FILTER_UNDEFINED {
		log.Printf("ERROR: filter %q undefined", filter)
		return
	}
	// Request
	req := &filterpb.FilterRequest{
		Filter: filterPB,
		Value:  value,
	}
	// DO REQUEST
	resStream, err := fc.service.Filter(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling filter RPC: %v", err)
	}
	// READ STREAMING
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// it reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading filter stream: %v", err)
		}
		log.Printf("RPC Filter response: %v", msg.GetFruit())
	}
}
