package client

import (
	"log"
)

func (*fruitClient) FilterCC(filter, value string) {
	log.Println("Filter Concurrent... starting")
	log.Printf("Request: filter %q, value %q", filter, value)
}
