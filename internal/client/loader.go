package client

import (
	"log"
)

func (*fruitClient) Loader() {
	log.Println("Requesting loader...")
}
