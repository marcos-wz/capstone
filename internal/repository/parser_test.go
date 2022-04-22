package repository

import (
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRepo_ParseFruitCSV(t *testing.T) {
	numFieldsFruit := reflect.TypeOf(entity.Fruit{}).NumField()
	var testCases = []struct {
		name      string
		record    []string
		numFields int
		resp      *basepb.Fruit
		errors    []error
	}{
		{
			"success, no error validations",
			[]string{
				"1", "TestFruit", "Testing fruit", "green", "kg", "5.50", "1", "1", "Mexico", "1642802058", "1647899658",
			},
			numFieldsFruit,
			&basepb.Fruit{
				Id: 1, Name: "TestFruit", Description: "Testing fruit", Color: "green",
				Unit: basepb.Unit_UNIT_KG, Price: 1, Stock: 1, CaducateDays: 1, Country: basepb.Country_COUNTRY_MEXICO,
				CreateTime: 1642802058, UpdateTime: 1647899658,
			},
			[]error{error(nil), error(nil), error(nil), error(nil), error(nil), error(nil), error(nil), error(nil), error(nil), error(nil)},
		},
	}

	// RUN TESTS --------------------------------

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fruit, errs := parseFruitCSV(tc.record, tc.numFields)
			assert.Equal(t, tc.resp, fruit)
			assert.Equal(t, tc.errors, errs)
			t.Logf("Fruit: %v", fruit)
			t.Logf("Errors: %v", errs)
		})
	}
}

func TestRepo_ParseUnit(t *testing.T) {
	var testCases = []struct {
		name     string
		unit     string
		response basepb.Unit
	}{
		{
			"Should return `LB` unit",
			"lb",
			basepb.Unit_UNIT_LB,
		},
		{
			"Should return `KG` unit",
			"kg",
			basepb.Unit_UNIT_KG,
		},
		{
			"Should return `UNDEFINED` unit",
			"fake-unit",
			basepb.Unit_UNIT_UNDEFINED,
		},
	}
	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := parseUnit(tc.unit)
			assert.Equal(t, tc.response, resp)
			t.Logf("Parsed unit %q : %v", tc.unit, resp)
		})
	}
}

func TestRepo_ParseCountry(t *testing.T) {
	var testCases = []struct {
		name     string
		country  string
		response basepb.Country
	}{
		{
			"Should return `MEXICO` country",
			"Mexico",
			basepb.Country_COUNTRY_MEXICO,
		},
		{
			"Should return `BRAZIL` country",
			"Brazil",
			basepb.Country_COUNTRY_BRAZIL,
		},
		{
			"Should return `CANADA` country",
			"Canada",
			basepb.Country_COUNTRY_CANADA,
		},
		{
			"Should return `USA` country",
			"USA",
			basepb.Country_COUNTRY_USA,
		},
		{
			"Should return `UNDEFINED` country",
			"fake-country",
			basepb.Country_COUNTRY_UNDEFINED,
		},
	}
	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := parseCountry(tc.country)
			assert.Equal(t, tc.response, resp)
			t.Logf("Parsed country %q : %v", tc.country, resp)
		})
	}
}

func TestRepo_ParseCurrency(t *testing.T) {
	var testCases = []struct {
		name     string
		currency string
		response basepb.Currency
	}{
		{
			"Should return `MNX` currency",
			"mxn",
			basepb.Currency_CURRENCY_MXN,
		},
		{
			"Should return `BRL` currency",
			"brl",
			basepb.Currency_CURRENCY_BRL,
		},
		{
			"Should return `CAD` currency",
			"cad",
			basepb.Currency_CURRENCY_CAD,
		},
		{
			"Should return `USD` currency",
			"usd",
			basepb.Currency_CURRENCY_USD,
		},
		{
			"Should return `UNDEFINED` currency",
			"fake-currency",
			basepb.Currency_CURRENCY_UNDEFINED,
		},
	}
	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := parseCurrency(tc.currency)
			assert.Equal(t, tc.response, resp)
			t.Logf("Parsed country %q : %v", tc.currency, resp)
		})
	}
}
