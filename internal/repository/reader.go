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
	var parserErrs string
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
	// indexRecord indicates the number of row record, it is used for parser errors
	indexRecord := 0
	for {
		record, err := csvReader.Read()
		// End of file, returns the parsed fruit records found in file
		if err == io.EOF {
			if parserErrs != "" {
				return list, fmt.Errorf("parser error: %v", parserErrs)
			}
			return list, nil
		}
		indexRecord++
		// Add parsed fruit to the list
		fruit, errs := rp.parseFruitRecord(record)
		// NOTE: How to treath parser errors ?
		// 		Maybe it should return an map[string]error: 1) "file", 2) "parser" ?
		errsStr, ok := rp.validateParserRecordErrors(errs)
		if errsStr != "" {
			msgErr := fmt.Sprintf("Record(%d): %v", indexRecord, errsStr)
			parserErrs += msgErr
			log.Printf("Parser ERROR: %v", msgErr)
			if !ok {
				// Log error to the logs server
				continue
			}
		}
		list = append(list, *fruit)
	}
}

// Formats errs array to string, indicating error field position and description
// if error on ID field, return false
// NOTE: Format json response ?
func (*readerRepo) validateParserRecordErrors(errs []error) (string, bool) {
	var description string
	valid := true
	totalErrs := 0
	if errs[0] != nil {
		log.Println("ERROR ID: ", errs[0])
		valid = false
	}
	for i, err := range errs {
		if err != nil {
			totalErrs++
			description += fmt.Sprintf("{ Field(%d): %v }", i, err)
		}
	}
	return description, valid
}

// Always returns a fruit instance and validates the value field
// if an error is found, the default value is set
// field error for each field, is returned by numeric/index array position
func (*readerRepo) parseFruitRecord(fields []string) (*entity.Fruit, []error) {
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
