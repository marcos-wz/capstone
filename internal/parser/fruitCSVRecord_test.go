package parser

import (
	"fmt"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_ParseFruitCSVRecord(t *testing.T) {
	var testCases = []struct {
		name   string
		record []string
		resp   *basepb.Fruit
		err    string
	}{
		{
			"Should return a valid parsed fruit, no error validations",
			[]string{"1", "Test Fruit", "Testing fruit", "green", "kg", "5.50", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			&basepb.Fruit{Id: 1, Name: "Test Fruit", Description: "Testing fruit", Color: "green", Unit: basepb.Unit_UNIT_KG, Currency: basepb.Currency_CURRENCY_MXN, Price: 5.50, Stock: 1, CaducateDays: 1, Country: basepb.Country_COUNTRY_MEXICO, CreateTime: 1642802058, UpdateTime: 1647899658},
			"<nil>",
		},
		// ID cases
		{
			"Should return ID error: invalid syntax",
			[]string{"s", "Test Fruit", "Testing fruit", "green", "kg", "5.50", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Id' Error:Field validation for 'Id' failed on the 'numeric' tag",
		},
		{
			"Should return ID error: zero value not allowed",
			[]string{"0", "Test Fruit", "Testing fruit", "green", "kg", "5.50", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Id' Error:Field validation for 'Id' failed on the 'ne' tag",
		},
		// NAME
		{
			"Should return NAME error: length must be 2 characters at least",
			[]string{"1", "n", "Testing fruit", "green", "kg", "5.50", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Name' Error:Field validation for 'Name' failed on the 'gt' tag",
		},
		// DESCRIPTION
		{
			"Should return DESCRIPTION error: field required",
			[]string{"1", "Fruit Test", "", "green", "kg", "5.50", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Description' Error:Field validation for 'Description' failed on the 'required' tag",
		},
		// COLOR
		{
			"Should return COLOR error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "", "kg", "5.50", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Color' Error:Field validation for 'Color' failed on the 'required' tag",
		},
		// UNIT
		{
			"Should return UNIT error: one of",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "fake-unit", "5.50", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag",
		},
		// PRICE
		{
			"Should return PRICE error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "0.00", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Price' Error:Field validation for 'Price' failed on the 'ne' tag",
		},
		{
			"Should return PRICE error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "0", "mxn", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Price' Error:Field validation for 'Price' failed on the 'ne' tag",
		},
		// CURRENCY
		{
			"Should return CURRENCY error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "10.00", "", "1", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Currency' Error:Field validation for 'Currency' failed on the 'oneof' tag",
		},
		// STOCK
		{
			"Should return STOCK error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "10.00", "mxn", "", "1", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Stock' Error:Field validation for 'Stock' failed on the 'required' tag",
		},
		// CADUCATE DAYS
		{
			"Should return CADUCATE DAYS error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "10.00", "mxn", "0", "", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.CaducateDays' Error:Field validation for 'CaducateDays' failed on the 'required' tag",
		},
		{
			"Should return CADUCATE DAYS error: zero value",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "10.00", "mxn", "0", "0", "Mexico", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.CaducateDays' Error:Field validation for 'CaducateDays' failed on the 'ne' tag",
		},
		// COUNTRY
		{
			"Should return COUNTRY error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "10.00", "mxn", "1", "1", "", "1642802058", "1647899658"},
			nil,
			"Key: 'FruitCSVRecord.Country' Error:Field validation for 'Country' failed on the 'oneof' tag",
		},
		// CREATE TIME
		{
			"Should return CREATE TIME error: field required",
			[]string{"1", "Fruit Test", "Testing fruit", "green", "kg", "10.00", "mxn", "1", "1", "Mexico", "", ""},
			nil,
			"Key: 'FruitCSVRecord.CreateTime' Error:Field validation for 'CreateTime' failed on the 'required' tag",
		},
	}

	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fruit, err := NewFruitParser().ParseFruitCSVRecord(tc.record)
			assert.Equal(t, tc.resp, fruit)
			assert.Equal(t, tc.err, fmt.Sprintf("%v", err))
		})
	}
}
