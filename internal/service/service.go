package service

import (
	"github.com/marcos-wz/capstone/internal/entity"
	pb "github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
)

// DOMAIN ***********************************************

type iFruitService interface {
	// GetFilteredFruits gets Filtered Fruits from the repository, refactor by filter. Repository error propagation support
	// Param: a valid Filter. Return: List of fruits and errors.
	GetFilteredFruits(filter *filterpb.FilterRequest) ([]*pb.Fruit, *entity.FruitFilterError)
}

type fruitRepo interface {
	ReadFruits() ([]*pb.Fruit, *entity.ReadFruitsError)
	FetchFruits() ([]*pb.Fruit, error)
	WriteFruits(fruits []*pb.Fruit) error
}

type fruitService struct {
	repo fruitRepo
}

func NewFruitService(repo fruitRepo) iFruitService {
	return &fruitService{repo}
}
