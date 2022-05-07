package service

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
)

var DebugLevel uint32

// FruitRepo dependency injection
type FruitRepo interface {
	ReadFruits() ([]*basepb.Fruit, error)
	FetchFruits() ([]*basepb.Fruit, error)
	WriteFruit(fruits *basepb.Fruit) error
}

type fruitService struct {
	repo FruitRepo
	fruitpb.UnimplementedFruitServiceServer
}

func NewFruitService(repo FruitRepo) *fruitService {
	return &fruitService{repo: repo}
}
