package repository

import (
	pb "github.com/marcos-wz/capstone/proto/basepb"
)

// DOMAIN ***********************************************

// flag for displaying debug messages
var Debug bool

type iFruitRepo interface {

	// ReadFruits fruits reader repository, read all fruit records from the
	// csv file, and guarantee csv data integrity.
	ReadFruits() ([]*pb.Fruit, error)

	// FetchFruits fetch data from external fruit
	FetchFruits() ([]*pb.Fruit, error)

	// WriteFruits Write new fruits records to the csv file
	WriteFruits(fruits []*pb.Fruit) error
}

type fruitRepo struct {
	filePath    string
	externalAPI string
}

func NewFruitRepo(file, url string) iFruitRepo {
	return &fruitRepo{
		filePath:    file,
		externalAPI: url,
	}
}
