package server

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
)

var Debug bool

// FruitService Mock dependency injection
type FruitService interface {
	FilterFruits(filter *filterpb.FilterRequest) ([]*basepb.Fruit, error)
}

type server struct {
	service FruitService
	fruitpb.UnimplementedFruitServiceServer
}

// NewFruitServer returns a setup server
func NewFruitServer(svc FruitService) *server {
	return &server{service: svc}
}
