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
			"../../data/test/csv/fruits-test-ok.csv",
			[]entity.Fruit{
				{ID: 1, Name: "Pera", Description: "Fruta tropical", Color: "green", Unit: "lb", Price: 0, Stock: 0, CaducateDays: 0, Country: "Canada", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 2, Name: "Manzana", Description: "Fruta tropical", Color: "red", Unit: "kg", Price: 0, Stock: 0, CaducateDays: 0, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 3, Name: "Platano", Description: "Fruta tropical", Color: "yellow", Unit: "kg", Price: 0, Stock: 0, CaducateDays: 0, Country: "Brazil", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 4, Name: "Mandarina", Description: "Fruta tropical", Color: "orange", Unit: "kg", Price: 0, Stock: 0, CaducateDays: 0, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 5, Name: "Naranja", Description: "Fruta tropical", Color: "yellow", Unit: "lb", Price: 0, Stock: 0, CaducateDays: 0, Country: "USA", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			},
			nil,
		},
		{
			"Should return error `no such file or directory`",
			"",
			nil,
			&entity.ReadFruitsError{
				Type:  "Repo.FileError",
				Error: errors.New("open : no such file or directory"),
			},
		},
		{
			"Should return partial fruits list, with parser error",
			"../../data/test/csv/fruits-test-error.csv",
			[]entity.Fruit{
				{ID: 0, Name: "TestFruit1", Description: "Fruta tropical", Color: "green", Unit: "lbs", Price: 0, Stock: 0, CaducateDays: 0, Country: "Canada", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 0, Name: "TestFruit1", Description: "Fruta tropical", Color: "green", Unit: "lbs", Price: 0, Stock: 0, CaducateDays: 0, Country: "Canada", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
				{ID: 0, Name: "", Description: "", Color: "", Unit: "", Price: 0, Stock: 0, CaducateDays: 0, Country: "", CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
			&entity.ReadFruitsError{
				Type:  "Repo.ParserError",
				Error: errors.New("reader repository, parse fruit errors found"),
				ParserErrors: []entity.ParseFruitRecordCSVError{
					{Record: 1, Errors: []entity.ParseFruitFieldCSVError{
						{Field: "ID", Value: "0", Validation: "required", Error: "Key: 'Fruit.ID' Error:Field validation for 'ID' failed on the 'required' tag"},
						{Field: "Unit", Value: "lbs", Validation: "oneof", Error: "Key: 'Fruit.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag"},
					}},
					{Record: 2, Errors: []entity.ParseFruitFieldCSVError{
						{Field: "ID", Value: "0", Validation: "required", Error: "Key: 'Fruit.ID' Error:Field validation for 'ID' failed on the 'required' tag"},
						{Field: "Unit", Value: "lbs", Validation: "oneof", Error: "Key: 'Fruit.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag"},
					}},
					{Record: 3, Errors: []entity.ParseFruitFieldCSVError{
						{Field: "ID", Value: "0", Validation: "required", Error: "Key: 'Fruit.ID' Error:Field validation for 'ID' failed on the 'required' tag"},
						{Field: "Name", Value: "", Validation: "required", Error: "Key: 'Fruit.Name' Error:Field validation for 'Name' failed on the 'required' tag"},
						{Field: "Color", Value: "", Validation: "required", Error: "Key: 'Fruit.Color' Error:Field validation for 'Color' failed on the 'required' tag"},
						{Field: "Unit", Value: "", Validation: "oneof", Error: "Key: 'Fruit.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag"},
						{Field: "Country", Value: "", Validation: "required", Error: "Key: 'Fruit.Country' Error:Field validation for 'Country' failed on the 'required' tag"},
						{Field: "CreatedAt", Value: "0001-01-01 00:00:00 +0000 UTC", Validation: "required", Error: "Key: 'Fruit.CreatedAt' Error:Field validation for 'CreatedAt' failed on the 'required' tag"},
					}},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewReaderRepo(tc.filePath)
			fruits, err := repo.ReadFruits()
			// Error
			if err != nil {
				assert.Equal(t, tc.err.Type, err.Type)
				assert.EqualError(t, err.Error, tc.err.Error.Error())
				assert.Len(t, err.ParserErrors, len(tc.err.ParserErrors))
				assert.EqualValues(t, tc.err.ParserErrors, err.ParserErrors)
			}
			// Response
			assert.Len(t, fruits, len(tc.response))
			assert.EqualValues(t, tc.response, fruits)
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
		err         *entity.ParseFruitFieldCSVError
	}{
		// ID
		{
			"success, no error validations",
			[]string{"1", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			nil,
		},
		{
			"Should return ID error tag: required validation",
			[]string{"", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 0, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "ID", Validation: "required",
				Error: "Key: 'Fruit.ID' Error:Field validation for 'ID' failed on the 'required' tag",
			},
		},
		// NAME
		{
			"Should return NAME error tag: required validation",
			[]string{"1", "", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "Name", Validation: "required",
				Error: "Key: 'Fruit.Name' Error:Field validation for 'Name' failed on the 'required' tag",
			},
		},
		{
			"Should return NAME error tag: only letters and numbers allowed",
			[]string{"1", "TestFruit1-", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit1-", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "Name", Validation: "alphanum",
				Error: "Key: 'Fruit.Name' Error:Field validation for 'Name' failed on the 'alphanum' tag",
			},
		},
		// COLOR
		{
			"Should return COLOR error tag: required validation",
			[]string{"1", "TestFruit", "Testing fruit", "", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "Color", Validation: "required",
				Error: "Key: 'Fruit.Color' Error:Field validation for 'Color' failed on the 'required' tag",
			},
		},
		{
			"Should return COLOR error tag: only letters allowed",
			[]string{"1", "TestFruit", "Testing fruit", "green1", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green1", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "Color", Validation: "alpha",
				Error: "Key: 'Fruit.Color' Error:Field validation for 'Color' failed on the 'alpha' tag",
			},
		},
		// UNIT
		{
			"Should return UNIT error tag: oneof validation",
			[]string{"1", "TestFruit", "Testing fruit", "purple", "badUnit", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "purple", Unit: "badUnit", Price: 1, Stock: 1, CaducateDays: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "Unit", Validation: "oneof",
				Error: "Key: 'Fruit.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag",
			},
		},
		// COUNTRY
		{
			"Should return COUNTRY error tag: required validation",
			[]string{"1", "TestFruit", "Testing fruit", "purple", "kg", "1.25", "5", "1", "", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "purple", Unit: "kg", Price: 1.25, Stock: 5, CaducateDays: 1, CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "Country", Validation: "required",
				Error: "Key: 'Fruit.Country' Error:Field validation for 'Country' failed on the 'required' tag",
			},
		},
		{
			"Should return COUNTRY error tag: only letters allowed",
			[]string{"1", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Country12", "2022-02-01T12:14:05-06:00"},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg", Price: 1, Stock: 1, CaducateDays: 1, Country: "Country12", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			&entity.ParseFruitFieldCSVError{
				Field: "Country", Validation: "alpha",
				Error: "Key: 'Fruit.Country' Error:Field validation for 'Country' failed on the 'alpha' tag",
			},
		},
		// CREATED AT
		{
			"Should return CREATED AT error tag: required validation",
			[]string{"1", "TestFruit", "Testing fruit", "purple", "kg", "1.25", "5", "1", "Mexico", ""},
			&entity.Fruit{ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "purple", Unit: "kg", Price: 1.25, Stock: 5, CaducateDays: 1, Country: "Mexico"},
			&entity.ParseFruitFieldCSVError{
				Field: "CreatedAt", Validation: "required",
				Error: "Key: 'Fruit.CreatedAt' Error:Field validation for 'CreatedAt' failed on the 'required' tag",
			},
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
					assert.Equal(t, tc.err.Field, errV.StructField())
					assert.Equal(t, tc.err.Validation, errV.Tag())
					assert.Equal(t, tc.err.Error, errV.Error())
				}
			}
			assert.Equal(t, tc.reponse, fruit)
		})
	}
}
