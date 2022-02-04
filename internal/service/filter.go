package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN ***********************************************

type iFilterService interface {
	// Get Filtered Fruits from the repository
	GetFilteredFruits(filter, value string) ([]entity.Fruit, error)
}

type readerRepo interface {
	ReadFruits() ([]entity.Fruit, error)
}

type filterService struct {
	repo readerRepo
}

func NewFilterService(repo readerRepo) iFilterService {
	return &filterService{repo}
}

// IMPLEMENTATION ***************************************

func (f *filterService) GetFilteredFruits(filter, value string) ([]entity.Fruit, error) {
	fruits, errRepo := f.repo.ReadFruits()
	// Repository errors evaluations
	if errRepo != nil {
		// Parser error propagation,
		// Return fruits with default values and/or ommited invalid records(lost data)
		if strings.HasPrefix(errRepo.Error(), "parser error: ") {
			filterdFruits, errFilter := f.filterFactory(fruits, filter, value)
			if errFilter != nil {
				return nil, errFilter
			}
			return filterdFruits, errRepo
		}
		// Default repository error
		return nil, errRepo
	}
	return f.filterFactory(fruits, filter, value)
}

// return fruits by filter, if not valid filter returns 0
func (f *filterService) filterFactory(fruits []entity.Fruit, filter, value string) ([]entity.Fruit, error) {
	switch filter {
	case "id":
		id, err := strconv.Atoi(value)
		if err != nil {
			err := fmt.Errorf("invalid ID filter(%v): %v", value, err)
			log.Println("ERROR SERVICE:", err)
			return nil, err
		}
		return f.filterByID(fruits, id), nil
	case "name":
		return f.filterByName(fruits, value), nil
	case "color":
		return f.filterByColor(fruits, value), nil
	case "country":
		return f.filterByCountry(fruits, value), nil
	case "all":
		return fruits, nil
	default:
		err := fmt.Errorf("undefined filter(%v): %v", filter, value)
		log.Println("ERROR Service:", err)
		return nil, err
	}
}

// Return filtered fruits records by ID
func (*filterService) filterByID(fruits []entity.Fruit, id int) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if fruit.ID == id {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

//  Return filtered fruits records by Name
func (*filterService) filterByName(fruits []entity.Fruit, name string) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Name, name) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

//  Return filtered fruits records by Color
func (*filterService) filterByColor(fruits []entity.Fruit, color string) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Color, color) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

//  Return filtered fruits records by Country
func (*filterService) filterByCountry(fruits []entity.Fruit, country string) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Country, country) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}
