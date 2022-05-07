package repository

import (
	"github.com/marcos-wz/capstone/proto/basepb"
)

// DebugLevel flag for verbose debugging messages
var DebugLevel uint32

type FruitRepo interface {

	// ReadFruits fruits reader repository, read all fruit records from the
	// csv file, and guarantee csv data integrity.
	ReadFruits() ([]*basepb.Fruit, error)

	// FetchFruits fetch data from external fruit
	FetchFruits() ([]*basepb.Fruit, error)

	// WriteFruit Write new fruit records to a csv file
	WriteFruit(fruit *basepb.Fruit) error
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
