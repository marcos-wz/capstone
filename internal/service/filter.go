package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN ***********************************************

type iFilter interface {
	// Get Filtered Fruits from the repository
	FilterFruits(filter, value string) ([]entity.Fruit, error)
}

// NOTE: reader or getter ?
type reader interface {
	ReadFruits() ([]entity.Fruit, error)
}

type filter struct {
	repo reader
}

func NewFilter(repo reader) iFilter {
	return &filter{repo}
}

// IMPLEMENTATION ***************************************

func (f *filter) FilterFruits(filter, value string) ([]entity.Fruit, error) {
	fruits, errRepo := f.repo.ReadFruits()
	// ********************
	// NOTE: fruits with parser errors
	// should return fruit valid fruit list, with default values and excluding invalid records?
	// should i create a parser ERROR to type validation, inteast of string ?
	if strings.Contains(fmt.Sprintf("%v", errRepo), "parser error:") {
		filterdFruits, errFilter := f.filterFactory(fruits, filter, value)
		if errFilter != nil {
			return nil, errFilter
		}
		return filterdFruits, errRepo
	}
	// ********************
	if errRepo != nil {
		return nil, errRepo
	}
	return f.filterFactory(fruits, filter, value)
}

// return fruits by filter, if not valid filter returns all fruits
func (f *filter) filterFactory(fruits []entity.Fruit, filter, value string) ([]entity.Fruit, error) {
	switch filter {
	case "id":
		id, err := strconv.Atoi(value)
		if err != nil {
			err := fmt.Errorf("invalid ID filter: %v - %v", value, err)
			log.Println("ERROR:", err)
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
		err := fmt.Errorf("invalid filter: %v - %v", filter, value)
		log.Println("ERROR:", err)
		return nil, err
	}
}

func (*filter) filterByID(fruits []entity.Fruit, id int) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if fruit.ID == id {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

func (*filter) filterByName(fruits []entity.Fruit, name string) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Name, name) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

func (*filter) filterByColor(fruits []entity.Fruit, color string) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Color, color) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

func (*filter) filterByCountry(fruits []entity.Fruit, country string) []entity.Fruit {
	filterdFruits := []entity.Fruit{}
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Country, country) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}
