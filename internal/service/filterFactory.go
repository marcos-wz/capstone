package service

import (
	"fmt"
	"github.com/marcos-wz/capstone/internal/parser"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"log"
	"strconv"
	"strings"
)

// return fruits by filter, if not valid filter returns an empty list
func filterFactory(fruits []*basepb.Fruit, filter *filterpb.FilterRequest) ([]*basepb.Fruit, error) {
	switch filter.Filter {
	case filterpb.FiltersAllowed_FILTER_ID:
		id, err := strconv.ParseUint(filter.Value, 10, 32)
		if err != nil {
			err := fmt.Errorf("invalid ID filter %q: %v", id, err)
			log.Println("SVC-ERROR:", err)
			return nil, err
		}
		return filterByID(fruits, uint32(id)), nil
	case filterpb.FiltersAllowed_FILTER_NAME:
		return filterByName(fruits, filter.Value), nil
	case filterpb.FiltersAllowed_FILTER_COLOR:
		return filterByColor(fruits, filter.Value), nil
	case filterpb.FiltersAllowed_FILTER_COUNTRY:
		country := parser.NewFruitParser().ParseCountry(filter.Value)
		return filterByCountry(fruits, country), nil
	default:
		err := fmt.Errorf("undefined filter(%v): %v", filter.Filter, filter.Value)
		log.Println("ERROR Service:", err)
		return nil, err
	}
}

// Return filtered fruits records by ID
func filterByID(fruits []*basepb.Fruit, id uint32) []*basepb.Fruit {
	if Debug {
		log.Printf("SVC: Filtering By ID %q...", id)
	}
	var filterdFruits []*basepb.Fruit
	for _, fruit := range fruits {
		if fruit.Id == id {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

//  Return filtered fruits records by Name
func filterByName(fruits []*basepb.Fruit, name string) []*basepb.Fruit {
	if Debug {
		log.Printf("SVC: Filtering By NAME %q...", name)
	}
	var filterdFruits []*basepb.Fruit
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Name, name) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

//  Return filtered fruits records by Color
func filterByColor(fruits []*basepb.Fruit, color string) []*basepb.Fruit {
	if Debug {
		log.Printf("SVC: Filtering by color %q...", color)
	}
	var filterdFruits []*basepb.Fruit
	for _, fruit := range fruits {
		if strings.EqualFold(fruit.Color, color) {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}

//  Return filtered fruits records by Country
func filterByCountry(fruits []*basepb.Fruit, country basepb.Country) []*basepb.Fruit {
	log.Printf("SVC: Filtering by country %q...", country)
	var filterdFruits []*basepb.Fruit
	for _, fruit := range fruits {
		if fruit.Country == country {
			filterdFruits = append(filterdFruits, fruit)
		}
	}
	return filterdFruits
}
