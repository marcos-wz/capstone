package client

import (
	"context"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"log"
	"time"
)

func (c *client) Filter(filter, value string) {
	log.Printf("Requesting Filter: %q, Value: %q ...", filter, value)
	// Input Validation
	filterAllowed := c.getAllowedFilter(filter)
	if filterAllowed == filterpb.FiltersAllowed_FILTER_UNDEFINED {
		log.Printf("ERROR: filter %q undefined", filter)
		return
	}
	// Service: contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.service.Filter(ctx, &filterpb.FilterRequest{
		Filter: filterAllowed,
		Value:  value,
	})
	if err != nil {
		log.Printf("ERROR: could not filter: %v", err)
		return
	}

	log.Printf("Resp Fruits: %s", r.GetFruit())
	//log.Printf("Resp Code: %d", r.GetCode())
	//log.Printf("Resp Fruits: %s", r.GetFruits())
}
