package repository

import (
	pb "github.com/marcos-wz/capstone/proto/basepb"
)

func (*fruitRepo) FetchFruits() ([]*pb.Fruit, error) {

	return []*pb.Fruit{}, nil
}
