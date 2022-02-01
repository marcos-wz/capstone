package repository

import "github.com/marcos-wz/capstone/internal/entity"

// DOMAIN ***********************************************

type iFetcher interface {
	// fetch data from external api
	FetchFruits() ([]entity.Fruit, error)
}

type fetcher struct {
	url string
}

func NewFetcher(url string) iFetcher {
	return &fetcher{url}
}

// IMPLEMENTATION ***********************************************

func (fetcher) FetchFruits() ([]entity.Fruit, error) {
	return nil, nil
}
