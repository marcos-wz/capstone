package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN ***********************************************

type iReaderRepo interface {
	// Read all fruit records from the csv file
	ReadFruits() ([]entity.Fruit, error)
}

type readerRepo struct {
	filePath string
}

func NewReaderRepo(file string) iReaderRepo {
	return &readerRepo{file}
}

// IMPLEMENTATION **************************************

func (rp *readerRepo) ReadFruits() ([]entity.Fruit, error) {
	// File
	f, err := os.Open(rp.filePath)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	defer f.Close()
	// CSV
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	list := []entity.Fruit{}
	numRecord := 0 // counter record is used for parser error description
	parserErrs := ""
	// Dynamic number of fruit entity struct
	numFruitFields := reflect.TypeOf(entity.Fruit{}).NumField()
	for {
		record, err := csvReader.Read()
		// End of file, returns the parsed fruit records found in file
		if err == io.EOF {
			// if parser error exists, returns the fruits parsed list, with default values set
			if parserErrs != "" {
				return list, fmt.Errorf("parser error: %v", parserErrs)
			}
			return list, nil
		}
		numRecord++
		// Add parsed fruit to the list
		fruit, errs := rp.parseFruitRecord(record, numFruitFields)
		if len(errs) > 0 {
			// Format to json
			parserErrs += rp.parserErrorsToJson(numRecord, errs)
			// Invalid ID records are ommited
			if _, ok := errs["id"]; ok {
				continue
			}
		}
		list = append(list, *fruit)
	}
}

// Always returns a fruit instance and validated fields values.
// If an error occurs or field not exists, the default type value is set.
// Returns a field value error array, the field is returning in json format
// Params:
//	- record : the string array record from csv file
//	- numFields : the number of fruit entity struct fields
func (*readerRepo) parseFruitRecord(record []string, numFields int) (*entity.Fruit, map[string]error) {
	fruit := &entity.Fruit{}
	errs := make(map[string]error)
	var err error
	// Initial values
	values := make([]string, numFields)
	copy(values, record)
	// VALIDATIONS
	// ID, must be integer and non-zero value
	fruit.ID, err = strconv.Atoi(values[0])
	if err != nil {
		errs["id"] = err
	} else {
		if fruit.ID == 0 {
			errs["id"] = errors.New("zero value error")
		}
	}
	// NAME
	fruit.Name = values[1]
	// DESCRIPTION
	fruit.Description = values[2]
	// COLOR
	fruit.Color = values[3]
	// UNIT
	fruit.Unit = values[4]
	// PRICE
	if fruit.Price, err = strconv.ParseFloat(values[5], 64); err != nil {
		errs["price"] = err
	}
	// STOCK
	if fruit.Stock, err = strconv.Atoi(values[6]); err != nil {
		errs["stock"] = err
	}
	// CADUCATE
	if fruit.Caducate, err = strconv.Atoi(values[7]); err != nil {
		errs["caducate"] = err
	}
	// COUNTRY
	fruit.Country = values[8]
	// CREATED AT
	if fruit.CreatedAt, err = time.Parse(time.RFC3339, values[9]); err != nil {
		errs["created_at"] = err
	}
	return fruit, errs
}

// Formats errs array to JSON string.
// It indicates field error and description
func (*readerRepo) parserErrorsToJson(index int, errs map[string]error) string {
	jsonResponse := fmt.Sprintf("{ %q : %d,[", "record", index)
	for field, err := range errs {
		jsonResponse += fmt.Sprintf("{ %q: %q },", field, err)
	}
	jsonResponse += "]}"
	return jsonResponse
}
