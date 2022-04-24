package repository

import (
	"encoding/csv"
	pb "github.com/marcos-wz/capstone/proto/basepb"
	"io"
	"log"
	"os"
)

func (rp *fruitRepo) ReadFruits() ([]*pb.Fruit, error) {
	if Debug {
		log.Println("ReadFruits repository starting...")
	}
	// File
	f, err := os.Open(rp.filePath)
	if err != nil {
		log.Printf("REPO-ERROR: read fruits: %v", err)
		return nil, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Printf("REPO-ERROR: closing file : %v", err)
		}
	}(f)

	// CSV
	csvReader := csv.NewReader(f)
	var fruits []*pb.Fruit
	// Load from record
	for {
		record, err := csvReader.Read()
		// End of file, returns the parsed fruit records found in file
		if err == io.EOF {
			return fruits, nil
		}
		// Add parsed fruit to the list
		fruit, err := parseFruitCSV(record)
		// NOTE: this is a lost data error, must log to a log server(Datadog)
		if err != nil {
			continue
		}
		fruits = append(fruits, fruit)
	}
}
