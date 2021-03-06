package service

import "github.com/marcos-wz/capstone/internal/entity"

// DOMAIN ***********************************************

type iLoaderService interface {
	// Add new fruits from external API in the repository, and returns total fruit inserted
	LoadAPIFruits() (int, error)
}

type fetcher interface {
	FetchFruits() ([]entity.Fruit, error)
}

type writer interface {
	WriteFruits(fruits []entity.Fruit) error
}

type loaderService struct {
	repoFetcher fetcher
	repoWriter  writer
}

func NewLoaderService(f fetcher, w writer) iLoaderService {
	return &loaderService{f, w}
}

// IMPLEMENTATION ***************************************

func (*loaderService) LoadAPIFruits() (int, error) {
	// var totalInserted int
	// var errCreate error
	// fruits, err := fs.repo.FetchFruits()
	// if err != nil {
	// 	return 0, err
	// }
	// for _, fruit := range fruits {
	// 	if err := fs.repo.WriteFruit(&fruit); err != nil {
	// 		errCreate = err
	// 		continue
	// 	}
	// 	totalInserted++
	// }
	// return totalInserted, errCreate
	return 0, nil
}
