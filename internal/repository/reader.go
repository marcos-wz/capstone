package repository

import (
	"encoding/csv"
	"encoding/json"
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
	parserRecordErrs := []entity.ParseCVSFruitRecordError{}
	// Dynamic number of fruit entity struct
	numFruitFields := reflect.TypeOf(entity.Fruit{}).NumField()
	for {
		record, err := csvReader.Read()
		// End of file, returns the parsed fruit records found in file
		if err == io.EOF {
			// if parser error exists, returns the fruits parsed list, with default values set and the error
			// returned error is JSON format, with "parser error: " preffix
			if len(parserRecordErrs) > 0 {
				parserErrorsJson, _ := json.Marshal(parserRecordErrs)
				return list, fmt.Errorf("parser error: %v", string(parserErrorsJson))
			}
			return list, nil
		}
		numRecord++
		// Add parsed fruit to the list
		fruit, errs := rp.parseFruitCSV(record, numFruitFields)
		if len(errs) > 0 {
			// Append the parsed record errors
			parserRecordErrs = append(parserRecordErrs, entity.ParseCVSFruitRecordError{
				Record: numRecord, Errors: errs,
			})

			// Validate required fields. If required field error exists, the record is ommited
			// NOTE: this is a lost data error, taking some actions should be important
			for _, err := range errs {
				if err.Required {
					log.Printf("REPO Parser required field(%v) %q : %v", err.Index, err.Field, err.Error)
					// Take some logging actions
					continue
				}
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
func (*readerRepo) parseFruitCSV(record []string, numFields int) (*entity.Fruit, []entity.ParseCSVFruitFieldError) {
	fruit := &entity.Fruit{}
	errs := []entity.ParseCSVFruitFieldError{}
	var err error
	// Initial values
	values := make([]string, numFields)
	copy(values, record)
	// VALIDATIONS
	// 0 - ID, must be integer and non-zero value
	fruit.ID, err = strconv.Atoi(values[0])
	if err != nil {
		// errs["id"] = err
		errs = append(errs, entity.ParseCSVFruitFieldError{
			Index: 0, Field: "ID", Error: err.Error(), Required: true,
		})
	} else {
		if fruit.ID == 0 {
			errs = append(errs, entity.ParseCSVFruitFieldError{
				Index: 0, Field: "ID", Error: "zero value error", Required: true,
			})
		}
	}
	// 1 - NAME
	fruit.Name = values[1]
	// 2 - DESCRIPTION
	fruit.Description = values[2]
	// 3 - COLOR
	fruit.Color = values[3]
	// 4 - UNIT
	fruit.Unit = values[4]
	// 5 - PRICE
	if fruit.Price, err = strconv.ParseFloat(values[5], 64); err != nil {
		errs = append(errs, entity.ParseCSVFruitFieldError{
			Index: 5, Field: "Price", Error: err.Error(),
		})
	}
	// 6 - STOCK
	if fruit.Stock, err = strconv.Atoi(values[6]); err != nil {
		errs = append(errs, entity.ParseCSVFruitFieldError{
			Index: 6, Field: "Stock", Error: err.Error(),
		})
	}
	// 7 - CADUCATE
	if fruit.Caducate, err = strconv.Atoi(values[7]); err != nil {
		errs = append(errs, entity.ParseCSVFruitFieldError{
			Index: 7, Field: "Caducate", Error: err.Error(),
		})
	}
	// 8 - COUNTRY
	fruit.Country = values[8]
	// 9 - CREATED AT
	if fruit.CreatedAt, err = time.Parse(time.RFC3339, values[9]); err != nil {
		errs = append(errs, entity.ParseCSVFruitFieldError{
			Index: 9, Field: "CreatedAt", Error: err.Error(),
		})
	}
	return fruit, errs
}
