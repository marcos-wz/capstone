package service

import (
	pb "github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"log"
)

func (f *fruitService) GetFilteredFruits(filter *filterpb.FilterRequest) ([]*pb.Fruit, error) {
	if Debug {
		log.Println("SVC: GetFilteredFruits: starting...")
	}
	// Repository
	fruits, err := f.repo.ReadFruits()
	if err != nil {
		return nil, err
	}
	// Filter Fruit List
	filteredFruits, err := filterFactory(fruits, filter)
	if err != nil {
		return nil, err
	}
	return filteredFruits, nil
}
