package repository

import (
	"github.com/marcos-wz/capstone/proto/basepb"
)

func (*fruitRepo) FetchFruits() ([]*basepb.Fruit, error) {

	return []*basepb.Fruit{}, nil
}
