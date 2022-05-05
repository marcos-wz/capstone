package repository

import "github.com/marcos-wz/capstone/internal/entity"

//	Error range: 10 - 19 repository errors
const (
	ErrUndefined     entity.ErrorType = 0
	errUndefinedDesc string           = "undefined error"

	// ErrFileCSV file error on csv files
	ErrFileCSV     entity.ErrorType = 10
	errFileCSVDesc string           = "csv file error"

	// ErrFileJSON file error on json files
	ErrFileJSON     entity.ErrorType = 11
	errFileJSONDesc string           = "json file error"

	// ErrParseFruitCSV parse fruit from csv records error
	ErrParseFruitCSV     entity.ErrorType = 12
	errParseFruitCSVDesc string           = "parse fruit error"
)
