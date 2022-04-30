package repository

import (
	pb "github.com/marcos-wz/capstone/proto/basepb"
)

// DOMAIN ***********************************************

// Debug flag for displaying debug messages
var Debug bool

type FruitRepo interface {

	// ReadFruits fruits reader repository, read all fruit records from the
	// csv file, and guarantee csv data integrity.
	ReadFruits() ([]*pb.Fruit, error)

	// FetchFruits fetch data from external fruit
	FetchFruits() ([]*pb.Fruit, error)

	// WriteFruit Write new fruit records to a csv file
	WriteFruit(fruit *pb.Fruit) error
}

type fruitRepo struct {
	filePath    string
	externalAPI string
}

func NewFruitRepo(file, url string) FruitRepo {
	return &fruitRepo{
		filePath:    file,
		externalAPI: url,
	}
}
