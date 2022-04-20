package server

import (
	"github.com/marcos-wz/capstone/proto/fruitpb"
)

type server struct {
	// Services
	//filterService   FilterService
	//loaderService   LoaderService
	//filterCCService FilterCCService

	fruitpb.UnimplementedFruitServiceServer
}

// Buider server
//func NewServer(filterSvc FilterService, loaderSvc LoaderService, filterCCSvc FilterCCService) *server {
func NewServer() *server {
	return &server{
		//filterService:   filterSvc,
		//loaderService:   loaderSvc,
		//filterCCService: filterCCSvc,
	}
}
