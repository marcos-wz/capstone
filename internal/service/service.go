package service

import (
	pb "github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
)

var Debug bool

type FruitService interface {

	// GetFilteredFruits gets Filtered Fruits from the repository, refactor by filter. Repository error propagation support
	// Param: a valid Filter. Return: List of fruits and errors.
	GetFilteredFruits(filter *filterpb.FilterRequest) ([]*pb.Fruit, error)
}

// FruitRepo Mock dependency injection
type FruitRepo interface {
	ReadFruits() ([]*pb.Fruit, error)
	FetchFruits() ([]*pb.Fruit, error)
	WriteFruits(fruits []*pb.Fruit) error
}

type fruitService struct {
	repo FruitRepo
}

func NewFruitService(repo FruitRepo) FruitService {
	return &fruitService{repo}
}
