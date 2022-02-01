package repository

import "github.com/marcos-wz/capstone/internal/entity"

// DOMAIN ***********************************************

type iFetcherRepo interface {
	// fetch data from external api
	FetchFruits() ([]entity.Fruit, error)
}

type fetcherRepo struct {
	url string
}

func NewFetcher(url string) iFetcherRepo {
	return &fetcherRepo{url}
}

// IMPLEMENTATION ***********************************************

func (*fetcherRepo) FetchFruits() ([]entity.Fruit, error) {
	return nil, nil
}
