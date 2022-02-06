package repository

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestReader_ReadFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		response []entity.Fruit
		err      *entity.ReadFruitsError
	}{
		{
			"Should return all fruits: no errors",
			"../../data/test/fruits-test-ok.csv",
			[]entity.Fruit{
				{ID: 1, Name: "Pera", Description: "Fruta tropical", Color: "green", Unit: "lbs", Price: 0, Stock: 0, CaducateDays: 0, Country: "Canada", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 2, Name: "Manzana", Description: "Fruta tropical", Color: "red", Unit: "kg", Price: 0, Stock: 0, CaducateDays: 0, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 3, Name: "Platano", Description: "Fruta tropical", Color: "yellow", Unit: "kg", Price: 0, Stock: 0, CaducateDays: 0, Country: "Brazil", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 4, Name: "Mandarina", Description: "Fruta tropical", Color: "orange", Unit: "kg", Price: 0, Stock: 0, CaducateDays: 0, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 5, Name: "Naranja", Description: "Fruta tropical", Color: "yellow", Unit: "lbs", Price: 0, Stock: 0, CaducateDays: 0, Country: "USA", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			},
			nil,
		},
		{
			"Should return error `no such file or directory`",
			"",
			nil,
			&entity.ReadFruitsError{Error: errors.New("open : no such file or directory")},
		},
		// {
		// 	"Should return valid fruit list, with parser error",
		// 	"../../data/test/fruits-test-error.csv",
		// 	[]entity.Fruit{
		// 		// Field ID
		// 		{Name: "TestFruit1", Description: "Fruta tropical", Color: "green", Unit: "lbs", Price: 0, Stock: 0, Caducate: 0, Country: "Canada", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
		// 		{ID: 0, Name: "TestFruit1", Description: "Fruta tropical", Color: "green", Unit: "lbs", Price: 0, Stock: 0, Caducate: 0, Country: "Canada", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
		// 		// All fields validations
		// 		{ID: 0},
		// 	},
		// 	&entity.ReadFruitsError{
		// 		Error: errors.New("cvs parser error"),
		// 		ParserError: []entity.CSVReaderParsedFruitError{
		// 			// Field ID empty validation
		// 			{Record: 1, Errors: []entity.CSVFruitFieldError{
		// 				{Index: 0, Field: "ID", Error: "strconv.Atoi: parsing \"\": invalid syntax", Required: true},
		// 			}},
		// 			// Field ID zero value validation
		// 			{Record: 2, Errors: []entity.CSVFruitFieldError{
		// 				{Index: 0, Field: "ID", Error: "zero value", Required: true},
		// 			}},
		// 			// All fields validation response
		// 			{Record: 3, Errors: []entity.CSVFruitFieldError{
		// 				{Index: 0, Field: "ID", Error: "strconv.Atoi: parsing \"-\": invalid syntax", Required: true},
		// 				{Index: 1, Field: "Name", Error: "empty value", Required: true},
		// 				{Index: 3, Field: "Color", Error: "empty value", Required: false},
		// 				{Index: 4, Field: "Unit", Error: "empty value", Required: false},
		// 				{Index: 5, Field: "Price", Error: "strconv.ParseFloat: parsing \"\": invalid syntax", Required: false},
		// 				{Index: 6, Field: "Stock", Error: "strconv.Atoi: parsing \"\": invalid syntax", Required: false},
		// 				{Index: 7, Field: "Caducate", Error: "strconv.Atoi: parsing \"\": invalid syntax", Required: false},
		// 				{Index: 9, Field: "CreatedAt", Error: "parsing time \"\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"2006\"", Required: false},
		// 			}},
		// 		},
		// 	},
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewReaderRepo(tc.filePath)
			fruits, err := repo.ReadFruits()
			// Errors
			if err != nil {
				assert.EqualError(t, err.Error, tc.err.Error.Error())
				assert.Equal(t, tc.err.ParserError, err.ParserError)
			}
			// Response
			assert.Equal(t, len(tc.response), len(fruits))
			assert.EqualValues(t, tc.response, fruits)
			// Debug ---------
			// if len(fruits) > 0 {
			// 	t.Log("Total fruits:", len(fruits))
			// 	for _, f := range fruits {
			// 		t.Logf("%+v", f)
			// 	}
			// }
		})
	}
}

func TestReader_ParseFruitRecord(t *testing.T) {
	numFieldsFruit := reflect.TypeOf(entity.Fruit{}).NumField()
	// Validator by struct TAG use cases
	var testCases = []struct {
		name        string
		recordParam []string
		reponse     *entity.Fruit
		errField    string
		errTag      string
		err         string
	}{
		// ID
		{
			"success, no error validations",
			[]string{"1", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"ID", "", "",
		},
		{
			"Should return ID error tag: required validation",
			[]string{"", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 0, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"ID", "required", "Key: 'Fruit.ID' Error:Field validation for 'ID' failed on the 'required' tag",
		},
		// NAME
		{
			"Should return NAME error tag: required validation",
			[]string{"1", "", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"Name", "required", "Key: 'Fruit.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		},
		{
			"Should return NAME error tag: only letters and numbers allowed",
			[]string{"1", "TestFruit1-", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit1-", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"Name", "alphanum", "Key: 'Fruit.Name' Error:Field validation for 'Name' failed on the 'alphanum' tag",
		},
		// COLOR
		{
			"Should return COLOR error tag: required validation",
			[]string{"1", "TestFruit", "Testing fruit", "", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"Color", "required", "Key: 'Fruit.Color' Error:Field validation for 'Color' failed on the 'required' tag",
		},
		{
			"Should return COLOR error tag: only letters allowed",
			[]string{"1", "TestFruit", "Testing fruit", "green1", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green1", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"Color", "alpha", "Key: 'Fruit.Color' Error:Field validation for 'Color' failed on the 'alpha' tag",
		},
		// UNIT
		{
			"Should return UNIT error tag: oneof validation",
			[]string{"1", "TestFruit", "Testing fruit", "purple", "badUnit", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "purple", Unit: "badUnit", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"Unit", "oneof", "Key: 'Fruit.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag",
		},
		// COUNTRY
		{
			"Should return COUNTRY error tag: required validation",
			[]string{"1", "TestFruit", "Testing fruit", "purple", "kg", "1.25", "5", "1", "", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "purple", Unit: "kg", Price: 1.25, Stock: 5, CaducateDays: 1, CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"Country", "required", "Key: 'Fruit.Country' Error:Field validation for 'Country' failed on the 'required' tag",
		},
		{
			"Should return COUNTRY error tag: only letters allowed",
			[]string{"1", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Country12", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Country12", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			"Country", "alpha", "Key: 'Fruit.Country' Error:Field validation for 'Country' failed on the 'alpha' tag",
		},
		// CREATED AT
		{
			"Should return CREATED AT error tag: required validation",
			[]string{"1", "TestFruit", "Testing fruit", "purple", "kg", "1.25", "5", "1", "Mexico", ""},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "purple", Unit: "kg", Price: 1.25, Stock: 5, CaducateDays: 1, Country: "Mexico"},
			"CreatedAt", "required", "Key: 'Fruit.CreatedAt' Error:Field validation for 'CreatedAt' failed on the 'required' tag",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &readerRepo{}
			fruit, err := r.parseFruitCSV(tc.recordParam, numFieldsFruit)
			if err != nil {
				errs := err.(validator.ValidationErrors)
				// One use case per tag validation, only one error validation should be returned
				assert.Len(t, errs, 1)
				for _, errV := range errs {
					// field, tag  and error validation must match
					assert.Equal(t, tc.errField, errV.StructField())
					assert.Equal(t, tc.errTag, errV.Tag())
					assert.Equal(t, tc.err, errV.Error())
				}
			}
			assert.Equal(t, tc.reponse, fruit)
			// t.Logf("Tested Fruit: %+v", fruit)
		})
	}
}
