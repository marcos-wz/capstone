package service

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
)

var Debug bool

type FruitService interface {

	// GetFilteredFruits gets Filtered Fruits from the repository, refactor by filter. Repository error propagation support
	// Param: a valid Filter. Return: List of fruits and errors.
	GetFilteredFruits(filter *filterpb.FilterRequest) ([]*basepb.Fruit, error)
}

// FruitRepo Mock dependency injection
type FruitRepo interface {
	ReadFruits() ([]*basepb.Fruit, error)
	FetchFruits() ([]*basepb.Fruit, error)
	WriteFruits(fruits []*basepb.Fruit) error
}

type fruitService struct {
	repo FruitRepo
}

func NewFruitService(repo FruitRepo) FruitService {
	return &fruitService{repo}
}
