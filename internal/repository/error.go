package repository

import "github.com/marcos-wz/capstone/internal/entity"

//	Error range: 10 - 19 repository errors
const (
	ErrUndefined entity.ErrorType = 0

	// ErrFileCSV file error on csv files
	ErrFileCSV entity.ErrorType = 10

	// ErrFileJSON file error on json files
	ErrFileJSON entity.ErrorType = 11

	// ErrParseFruitCSV parse fruit from csv records error
	ErrParseFruitCSV           entity.ErrorType = 12
	ErrParseFruitCSVIndex      entity.ErrorType = 13
	ErrParseFruitCSVValidation entity.ErrorType = 14
)

// ErrDesc holds repository errors description
var ErrDesc = map[entity.ErrorType]string{
	ErrUndefined:               "undefined error",
	ErrFileCSV:                 "csv file error",
	ErrFileJSON:                "json file error",
	ErrParseFruitCSV:           "parse fruit error",
	ErrParseFruitCSVIndex:      "parse fruit csv record undefined index error",
	ErrParseFruitCSVValidation: "parse fruit csv record validation error",
}
