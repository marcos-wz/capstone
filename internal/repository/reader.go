package repository

import (
	"github.com/marcos-wz/capstone/internal/entity"
	pb "github.com/marcos-wz/capstone/proto/basepb"
	"log"
)

// READ FRUITS IMPLEMENTATION **************************************

func (rp *fruitRepo) ReadFruits() ([]*pb.Fruit, *entity.ReadFruitsError) {
	log.Println("ReadFruits repository starting...")
	// File
	//f, err := os.Open(rp.filePath)
	//if err != nil {
	//	log.Println("ERROR Reader Repo:", err)
	//	return nil, &entity.ReadFruitsError{
	//		Type:  "Repo.FileError",
	//		Error: err,
	//	}
	//}
	//defer f.Close()
	//// CSV
	//csvReader := csv.NewReader(f)
	//csvReader.FieldsPerRecord = -1
	//fruits := []*proto.Fruit{}
	//parserErrors := []entity.ParseFruitRecordCSVError{}
	//// counter record is used for parser error description
	//numRecord := 0
	//// Set dynamic number of fields, based on Fruit entity
	//numRecordFields := reflect.TypeOf(proto.Fruit{}).NumField()
	//for {
	//	record, err := csvReader.Read()
	//	// End of file, returns the parsed fruit records found in file
	//	if err == io.EOF {
	//		// if parser validation errors exists, returns the partial parsed fruits, with default values set and the field error validations
	//		// NOTE: this is a lost data error, taking some actions should be important
	//		if len(parserErrors) > 0 {
	//			return fruits, &entity.ReadFruitsError{
	//				Type:         "Repo.ParserError",
	//				Error:        errors.New("reader repository, parse fruit errors found"),
	//				ParserErrors: parserErrors,
	//			}
	//		}
	//		return fruits, nil
	//	}
	//	numRecord++
	//	// Add parsed fruit to the list
	//	fruit, err := rp.parseFruitCSV(record, numRecordFields)
	//	if err != nil {
	//		validationErrors := err.(validator.ValidationErrors)
	//		// Parser validation errors to ParseFruitFieldCSVError type
	//		fieldErrs := []entity.ParseFruitFieldCSVError{}
	//		for _, vErr := range validationErrors {
	//			fieldErrs = append(fieldErrs, entity.ParseFruitFieldCSVError{
	//				Field:      vErr.Field(),
	//				Value:      fmt.Sprintf("%v", vErr.Value()),
	//				Validation: vErr.ActualTag(),
	//				Error:      vErr.Error(),
	//			})
	//		}
	//		// Append the parsed record validations errors
	//		parserErrors = append(parserErrors, entity.ParseFruitRecordCSVError{
	//			Record: numRecord, Errors: fieldErrs,
	//		})
	//		// Validate required fields. If required field error exists, the record is ommited
	//		// NOTE: this is a lost data error, taking some actions should be important
	//		for _, vErr := range validationErrors {
	//			if vErr.StructField() == "ID" || vErr.StructField() == "Name" {
	//				log.Printf("ERROR Repo: Invalid record(lost data): field: %q, value: %v, error:  %v ", vErr.StructField(), vErr.Value(), vErr.Error())
	//				continue
	//			}
	//		}
	//	}
	//	fruits = append(fruits, fruit)
	//}
	return []*pb.Fruit{}, nil
}
