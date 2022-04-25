package server

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
)

var Debug bool

// FruitService Mock dependency injection
type FruitService interface {
	GetFilteredFruits(filter *filterpb.FilterRequest) ([]*basepb.Fruit, error)
}

type server struct {
	service FruitService
	fruitpb.UnimplementedFruitServiceServer
}

// NewServer returns a setup server
func NewServer(svc FruitService) *server {
	return &server{service: svc}
}
