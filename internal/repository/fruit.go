package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN ***********************************************

type iFruitRepo interface {
	// Read a fruit record by ID field from the csv file
	ReadFruitByID(id int) (*entity.Fruit, error)
	// Write new fruit intances to the csv file
	WriteFruit(fruit *entity.Fruit) error
	// Read all fruit records from the csv file
	ReadAllFruits() (*entity.Fruits, error)
}

type fruitRepo struct {
	csvFile string
}

func NewFruitRepo(file string) iFruitRepo {
	return &fruitRepo{file}
}

// IMPLEMENTATIONS **************************************

func (fr *fruitRepo) ReadFruitByID(id int) (*entity.Fruit, error) {
	f, err := os.Open(fr.csvFile)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		// End of file validation, fruit not found
		if err == io.EOF {
			log.Printf("ERROR: ID fruit %d not found - %v", id, err)
			return nil, err
		}
		// Parse fruit
		fruit, _ := fr.parseFruit(record)
		// Validate fruit ID
		if fruit.ID == id {
			return fruit, nil
		}
	}
}

func (fr *fruitRepo) WriteFruit(fruit *entity.Fruit) error {
	f, err := os.OpenFile(fr.csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	defer f.Close()
	// Parse fruit struct to string array
	numFields := 10
	// NOTE: check if dynamic legth make sense
	// t := reflect.TypeOf(*fruit)
	// numFields := t.NumField()

	fruitRecord := make([]string, numFields)
	fruitRecord[0] = strconv.Itoa(fruit.ID)
	fruitRecord[1] = fruit.Name
	fruitRecord[2] = fruit.Description
	fruitRecord[3] = fruit.Color
	fruitRecord[4] = fruit.Unit
	// fruitRecord[5] = fmt.Sprintf("%f", fruit.Price)
	fruitRecord[5] = fmt.Sprintf("%.2f", fruit.Price)
	fruitRecord[6] = strconv.Itoa(fruit.Stock)
	fruitRecord[7] = strconv.Itoa(fruit.Caducate)
	fruitRecord[8] = fruit.Country
	fruitRecord[9] = fruit.CreateAt.Format(time.RFC3339)

	log.Printf("Parsed fruit record: %+v", fruitRecord)
	// Write record
	csvWriter := csv.NewWriter(f)
	if err := csvWriter.Write(fruitRecord); err != nil {
		log.Println("ERROR:", err)
		return err
	}
	csvWriter.Flush()
	return nil
}

func (fr *fruitRepo) ReadAllFruits() (*entity.Fruits, error) {
	// File
	f, err := os.Open(fr.csvFile)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	defer f.Close()
	// CSV
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	list := entity.Fruits{}
	for {
		record, err := csvReader.Read()
		// End of file, returns the parsed fruit records found in file
		if err == io.EOF {
			return &list, nil
		}
		// Add parsed fruit to the list
		fruit, _ := fr.parseFruit(record)
		list = append(list, *fruit)
	}
}

// METHODS ********************************************

// Always returns a fruit instance and validates the value field
// if an error is found, the default value is set
// field error for each field, is returned by numeric/index array position
func (*fruitRepo) parseFruit(fields []string) (*entity.Fruit, []error) {
	fruit := &entity.Fruit{}
	numFields := 10
	// NOTE: check if dynamic legth make sense
	// t := reflect.TypeOf(*fruit)
	// numFields := t.NumField()
	errs := make([]error, numFields)
	values := make([]string, numFields)
	copy(values, fields)

	// ID, NOTE: check if it is a required field, if so must return error ?, ¿ ID = 0, could be duplicated ?
	fruit.ID, errs[0] = strconv.Atoi(values[0])
	// NAME
	fruit.Name = values[1]
	// DESCRIPTION
	fruit.Description = values[2]
	// COLOR
	fruit.Color = values[3]
	// UNIT
	fruit.Unit = values[4]
	// PRICE
	fruit.Price, errs[5] = strconv.ParseFloat(values[5], 64)
	// STOCK
	fruit.Stock, errs[6] = strconv.Atoi(values[6])
	// CADUCATE
	fruit.Caducate, errs[7] = strconv.Atoi(values[7])
	// COUNTRY
	fruit.Country = values[8]
	// CREATED AT
	fruit.CreateAt, errs[9] = time.Parse(time.RFC3339, values[9])
	return fruit, errs
}
