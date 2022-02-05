package repository

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN ***********************************************

type iReaderRepo interface {
	// Read all fruit records from the csv file
	ReadFruits() ([]entity.Fruit, *entity.ReadFruitsError)
}

type readerRepo struct {
	filePath string
}

func NewReaderRepo(file string) iReaderRepo {
	return &readerRepo{file}
}

// IMPLEMENTATION **************************************

func (rp *readerRepo) ReadFruits() ([]entity.Fruit, *entity.ReadFruitsError) {
	// File
	f, err := os.Open(rp.filePath)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, &entity.ReadFruitsError{Error: err}
	}
	defer f.Close()
	// CSV
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	list := []entity.Fruit{}
	numRecord := 0 // counter record is used for parser error description
	parserRecordErrs := []entity.CSVReaderParsedFruitError{}
	// Dynamic number of fruit entity struct
	numRecordFields := reflect.TypeOf(entity.Fruit{}).NumField()
	for {
		record, err := csvReader.Read()
		// End of file, returns the parsed fruit records found in file
		if err == io.EOF {
			// if parser error exists, returns the partial parsed fruits, with default values set and the field error
			// with "parser error:"
			if len(parserRecordErrs) > 0 {
				return list, &entity.ReadFruitsError{
					Error:       errors.New("cvs parser error"),
					ParserError: parserRecordErrs,
				}
			}
			return list, nil
		}
		numRecord++
		// Add parsed fruit to the list
		fruit, err := rp.parseFruitCSV(record, numRecordFields)
		if err != nil {
			// Append the parsed record errors
			parserRecordErrs = append(parserRecordErrs,
				entity.CSVReaderParsedFruitError{Record: numRecord, Error: err},
			)
			// Validate required fields. If required field error exists, the record is ommited
			// NOTE: this is a lost data error, taking some actions should be important
			// for _, err := range errs {
			// 	if err.Required {
			// 		log.Printf("REPO Parser required field: %v - %q: %v", err.Index, err.Field, err.Error)
			// 		// Take some logging actions
			// 		continue
			// 	}
			// }
		}
		list = append(list, *fruit)
	}
}

// Input data method. Always returns a valid fruit.
// If an error occurs, the default type value is set.
// Params:
//	- record : the string array csv library format record from file
//	- numFields : the number of fruit entity struct fields
func (*readerRepo) parseFruitCSV(record []string, numFields int) (*entity.Fruit, error) {
	// Initial values
	fruit := &entity.Fruit{}
	values := make([]string, numFields)
	copy(values, record)
	// Fruit Instance
	// 0 - ID
	fruit.ID, _ = strconv.Atoi(values[0])
	// 1 - NAME
	fruit.Name = values[1]
	// 2 - DESCRIPTION
	fruit.Description = values[2]
	// 3 - COLOR
	fruit.Color = values[3]
	// 4 - UNIT
	fruit.Unit = values[4]
	// 5 - PRICE
	fruit.Price, _ = strconv.ParseFloat(values[5], 64)
	// 6 - STOCK
	fruit.Stock, _ = strconv.Atoi(values[6])
	// 7 - CADUCATE
	fruit.Caducate, _ = strconv.Atoi(values[7])
	// 8 - COUNTRY
	fruit.Country = values[8]
	// 9 - CREATED AT
	fruit.CreatedAt, _ = time.Parse(time.RFC3339, values[9])

	return fruit, validator.New().Struct(fruit)
}
