package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN ***************************************************************************

type iFilterService interface {
	// Get Filtered Fruits from the repository, refactor by filter. Repository error propagation support
	// Param: a valid Filter. Return: List of fruits and errors.
	GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, *entity.FruitFilterError)
}

type ReaderRepo interface {
	ReadFruits() ([]entity.Fruit, *entity.ReadFruitsError)
}

type filterService struct {
	repo ReaderRepo
}

func NewFilterService(repo ReaderRepo) iFilterService {
	return &filterService{repo}
}

// IMPLEMENTATION *******************************************************************

func (f *filterService) GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, *entity.FruitFilterError) {
	fruits, errRepo := f.repo.ReadFruits()
	// Repository error propagation
	if errRepo != nil {
		// Repository parser error propagation, returns partial fruit list, with default values
		if errRepo.Type == "Repo.ParserError" {
			filterdFruits, err := f.filterFactory(fruits, filter)
			if err != nil {
				return nil, &entity.FruitFilterError{
					Type:  "Service.FilterError",
					Error: err,
				}
			}
			// returns filtered partial fruits list, and repository parsed fruit errors
			return filterdFruits, &entity.FruitFilterError{
				Type:         errRepo.Type,
				Error:        errRepo.Error,
				ParserErrors: errRepo.ParserErrors,
			}
		}
		// Default repository error propagation
		return nil, &entity.FruitFilterError{
			Type:  errRepo.Type,
			Error: errRepo.Error,
		}
	}
	// Filter Fruit List
	filteredFruits, err := f.filterFactory(fruits, filter)
	if err != nil {
		return nil, &entity.FruitFilterError{
			Type:  "Service.FilterError",
			Error: err,
		}
	}
	return filteredFruits, nil
}

// return fruits by filter, if not valid filter returns an empty list
func (f *filterService) filterFactory(fruits []entity.Fruit, filter *entity.FruitsFilterParams) ([]entity.Fruit, error) {
	switch filter.Filter {
	case "id":
		id, err := strconv.Atoi(filter.Value)
		if err != nil {
			err := fmt.Errorf("invalid ID filter(%v): %v", filter.Value, err)
			log.Println("ERROR SERVICE:", err)
			return nil, err
		}
		return f.filterByID(fruits, id), nil
	case "name":
		return f.filterByName(fruits, filter.Value), nil
	case "color":
		return f.filterByColor(fruits, filter.Value), nil
	case "country":
		return f.filterByCountry(fruits, filter.Value), nil
	default:
		err := fmt.Errorf("undefined filter(%v): %v", filter.Filter, filter.Value)
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
