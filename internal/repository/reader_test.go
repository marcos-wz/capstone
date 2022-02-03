package repository

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestReader_ReadFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		response []entity.Fruit
		err      string
	}{
		{
			"Should return valid fruit list",
			"../../data/test/fruits-test-ok.csv",
			[]entity.Fruit{},
			"<nil>",
		},
		{
			"Should return error `no such file or directory`",
			"",
			nil,
			"open : no such file or directory",
		},
		{
			"Should return valid fruit list, with parser error",
			"../../data/test/fruits-test-error.csv",
			[]entity.Fruit{},
			`parser error: [{"record":4,"errors":[{"index":0,"field":"ID","error":"strconv.Atoi: parsing \"-\": invalid syntax","Required":true},{"index":5,"field":"Price","error":"strconv.ParseFloat: parsing \"\": invalid syntax","Required":false},{"index":6,"field":"Stock","error":"strconv.Atoi: parsing \"\": invalid syntax","Required":false},{"index":7,"field":"Caducate","error":"strconv.Atoi: parsing \"\": invalid syntax","Required":false},{"index":9,"field":"CreatedAt","error":"parsing time \"\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"2006\"","Required":false}]}]`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewReaderRepo(tc.filePath)
			fruits, err := repo.ReadFruits()
			assert.Equal(t, tc.err, fmt.Sprintf("%v", err))
			if len(fruits) > 0 {
				t.Log("Total fruits:", len(fruits))
				for _, f := range fruits {
					t.Logf("%+v", f)
				}
			}
		})
	}
}

func TestReader_ParseFruitRecord(t *testing.T) {
	var testCases = []struct {
		name           string
		recordParam    []string
		numFieldsParam int
		reponse        *entity.Fruit
		errs           []entity.ParseCSVFruitFieldError
	}{
		{
			"success, no errors",
			[]string{"1", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			reflect.TypeOf(entity.Fruit{}).NumField(),
			&entity.Fruit{
				ID: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg",
				Price: 1, Stock: 1, Caducate: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local)},
			[]entity.ParseCSVFruitFieldError{},
		},
		{
			"Should return ID error, strconv.Atoi: parsing invalid syntax",
			[]string{"s", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			reflect.TypeOf(entity.Fruit{}).NumField(),
			&entity.Fruit{
				ID: 0, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg",
				Price: 1, Stock: 1, Caducate: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local),
			},
			[]entity.ParseCSVFruitFieldError{
				{Index: 0, Field: "ID", Error: "strconv.Atoi: parsing \"s\": invalid syntax"},
			},
		},
		{
			"Should return ID error, zero value",
			[]string{"0", "TestFruit", "Testing fruit", "green", "kg", "1", "1", "1", "Mexico", "2022-02-01T12:14:05-06:00"},
			reflect.TypeOf(entity.Fruit{}).NumField(),
			&entity.Fruit{
				ID: 0, Name: "TestFruit", Description: "Testing fruit", Color: "green", Unit: "kg",
				Price: 1, Stock: 1, Caducate: 1, Country: "Mexico", CreatedAt: time.Date(2022, time.February, 1, 12, 14, 5, 0, time.Local),
			},
			[]entity.ParseCSVFruitFieldError{
				{Index: 0, Field: "ID", Error: "zero value error"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &readerRepo{}
			fruit, errs := r.parseFruitCSV(tc.recordParam, tc.numFieldsParam)
			assert.Equal(t, len(tc.errs), len(errs))
			if len(errs) > 0 {
				for i, e := range errs {
					assert.Equal(t, tc.errs[i].Error, e.Error)
				}
			} else {
				assert.NotEmpty(t, fruit)
			}
			assert.Equal(t, tc.reponse, fruit)
			// Debug ---------
			t.Log("Parser errors:", errs)
			t.Log("Parsed fruit:", fruit)
		})
	}
}
