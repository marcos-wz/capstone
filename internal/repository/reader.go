package repository

import (
	"encoding/csv"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/marcos-wz/capstone/proto/basepb"
	"io"
	"log"
	"os"
)

func (rp *fruitRepo) ReadFruits() ([]*basepb.Fruit, error) {
	if DebugLevel >= 1 {
		log.Println("ReadFruits repository starting...")
	}
	// File
	f, err := os.Open(rp.filePath)
	if err != nil {
		log.Printf("REPO-ERROR: read fruits: %v", err)
		return nil, &entity.FruitError{Type: ErrFileCSV, Desc: ErrDesc[ErrFileCSV], Err: err}
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Printf("REPO-FATAL-ERROR: closing file : %v", err)
		}
	}(f)

	// CSV
	csvReader := csv.NewReader(f)
	var fruits []*basepb.Fruit
	for {
		record, err := csvReader.Read()
		// End of file, returns the parsed fruit records found in file
		if err == io.EOF {
			return fruits, nil
		}
		// Add parsed fruit to the list
		fruit, err := ParseFruitRecord(record)
		if err != nil {
			return nil, err
		}
		fruits = append(fruits, fruit)
	}
}
