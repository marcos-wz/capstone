package server

import (
	"github.com/marcos-wz/capstone/internal/entity"
	pb "github.com/marcos-wz/capstone/internal/fruitpb"
)

type FilterService interface {
	GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, *entity.FruitFilterError)
}

type LoaderService interface {
}

type FilterCCService interface {
}

// Buider server
func NewServerServices(filterSvc FilterService, loaderSvc LoaderService, filterCCSvc FilterCCService) *server {
	return &server{
		filterService:   filterSvc,
		loaderService:   loaderSvc,
		filterCCService: filterCCSvc,
	}
}

type server struct {

	// Services
	filterService   FilterService
	loaderService   LoaderService
	filterCCService FilterCCService

	pb.UnimplementedFruitServiceServer
}
