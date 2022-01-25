package repository

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/marcos-wz/capstone/internal/entity"
)

// CSV file path
var csvFilePath = "/go/src/capstone/data/fruits.csv"

// DOMAIN / CONTRACT ***********************************

type FruitRepo interface {
	// Get fruit by ID
	ReadFruitByID(id int) (*entity.Fruit, error)
	// Add new fruit in the csv file
	CreateFruit(fruit *entity.Fruit) error
	// Returns full list of fruits from csv file
	ListFruits() (*entity.Fruits, error)
}

type fruitRepo struct{}

func NewFruitRepo() FruitRepo {
	return &fruitRepo{}
}

// IMPLEMENTATION **************************************

func (fruitRepo) parseFruit(fields []string) (*entity.Fruit, error) {
	log.Printf("Parsing record(%d): %+v", len(fields), fields)

	// validate entity number fields

	fruit := &entity.Fruit{}
	return fruit, nil
}

func (fr fruitRepo) ReadFruitByID(id int) (*entity.Fruit, error) {
	f, err := os.Open(csvFilePath)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		// Error end of file validation, fruit not found
		if err == io.EOF {
			log.Println("ERROR: fruit not found - ", err)
			return nil, err
		}
		if err != nil {
			log.Printf("ERROR: %v - %v", err, record)
			continue
		}
		// Parse fruit
		fruit, err := fr.parseFruit(record)
		if err != nil {
			return nil, err
		}
		// Validate fruit ID
		if fruit.ID == id {
			return fruit, nil
		}
	}
}

func (fruitRepo) CreateFruit(fruit *entity.Fruit) error {
	return nil
}

func (fruitRepo) ListFruits() (*entity.Fruits, error) {
	return nil, nil
}
