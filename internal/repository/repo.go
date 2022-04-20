package repository

//
//import (
//	"github.com/marcos-wz/capstone/internal/entity"
//	pb "github.com/marcos-wz/capstone/internal/fruit"
//)
//
//// DOMAIN ***********************************************
//
//type iFruitRepo interface {
//
//	// ReadFruits fruits reader repository, read all fruit records from the
//	// csv file, and guarantee csv data integrity.
//	ReadFruits() ([]*pb.Fruit, *entity.ReadFruitsError)
//
//	// FetchFruits fetch data from external fruit
//	FetchFruits() ([]*pb.Fruit, error)
//
//	// WriteFruits Write new fruits records to the csv file
//	WriteFruits(fruits []*pb.Fruit) error
//}
//
//type fruitRepo struct {
//	filePath    string
//	externalAPI string
//}
//
//func NewFruitRepo() iFruitRepo {
//	return &fruitRepo{}
//}
