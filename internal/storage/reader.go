package storage

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN  *******************************************

type Reader interface {
	// Returns array of fruits from the csv file
	ReadFruits() (entity.Fruits, error)
}

type reader struct{}

func NewReader() Reader {
	return &reader{}
}

// IMPLEMENTATION **************************************

func (reader) ReadFruits() (entity.Fruits, error) {
	f, err := os.Open(csvFilePath)
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
		fruits = append(fruits, *parseFruit(record))
	}
}
