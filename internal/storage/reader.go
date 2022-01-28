package storage

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN **********************************************

type iReader interface {
	// Reads fruit records from csv file and returns list of fruits
	ReadFruits() (entity.Fruits, error)
}

type reader struct {
	filePath string
}

func NewReader(file string) iReader {
	return &reader{file}
}

// IMPLEMENTATION **************************************

func (r *reader) ReadFruits() (entity.Fruits, error) {
	f, err := os.Open(r.filePath)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	fruits := entity.Fruits{}
	for {
		record, err := csvReader.Read()
		// End of file validation
		if err == io.EOF {
			return fruits, nil
		}
		if err != nil {
			log.Printf("WARNING: %v - %v", err, record)
		}
		// Append valid fruit to the response
		fruits = append(fruits, *r.parseFruit(record))
	}
}

// Returns a fruit instance, if some field convertion get error, set default data type value
func (*reader) parseFruit(fields []string) *entity.Fruit {
	var err error
	fruit := &entity.Fruit{}

	// ID
	fruit.ID, err = strconv.Atoi(fields[0])
	if err != nil {
		log.Println("ERROR Parsing ID:", err)
	}
	// NAME
	if len(fields) >= 2 {
		fruit.Name = fields[1]
	}
	// DESCRIPTION
	if len(fields) >= 3 {
		fruit.Description = fields[2]
	}
	// COLOR
	if len(fields) >= 4 {
		fruit.Color = fields[3]
	}
	// UNIT
	if len(fields) >= 5 {
		fruit.Unit = fields[4]
	}
	// PRICE
	if len(fields) >= 6 {
		fruit.Price, err = strconv.ParseFloat(fields[5], 64)
		if err != nil {
			log.Println("ERROR Parsing Price:", err)
		}
	}
	// STOCK
	if len(fields) >= 7 {
		fruit.Stock, err = strconv.Atoi(fields[6])
		if err != nil {
			log.Println("ERROR Parsing Stock:", err)
		}
	}
	// CADUCATE
	if len(fields) >= 8 {
		fruit.Caducate, err = strconv.Atoi(fields[7])
		if err != nil {
			log.Println("ERROR Parsing Caducate:", err)
		}
	}
	// COUNTRY
	if len(fields) >= 9 {
		fruit.Country = fields[8]
	}
	// CREATED AT
	if len(fields) >= 10 {
		fruit.CreateAt, err = time.Parse(time.RFC3339, fields[9])
		if err != nil {
			log.Println("ERROR Parsing CreatedAt:", err)
		}
	}
	return fruit
}
