package server

import (
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
)

type server struct {
	service fruitService
	fruitpb.UnimplementedFruitServiceServer
}

type fruitService interface {
	GetFilteredFruits(filter *filterpb.FilterRequest) ([]*basepb.Fruit, *entity.FruitFilterError)
}

// NewServer returns a setup server
func NewServer(svc fruitService) *server {
	return &server{service: svc}
}
