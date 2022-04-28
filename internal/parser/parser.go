package parser

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
)

var Debug bool

type FruitParser interface {

	// ParseFilter turns a valid protobuf filter
	ParseFilter(filter string) filterpb.FiltersAllowed

	// ParseUnit returns a parsed protobuf unit
	ParseUnit(unit string) basepb.Unit

	// ParseCurrency parse a currency and returns a currency prototype
	ParseCurrency(currency string) basepb.Currency

	// ParseCountry parse a country and returns a unit prototype
	ParseCountry(country string) basepb.Country

	// ParseFruitCSVRecord guarantee csv data integrity. It parses from csv records to fruit instance.
	// Record parameters is a string array from csv file. Returns a parsed fruit instance or error.
	ParseFruitCSVRecord(record []string) (*basepb.Fruit, error)

	// ParseFruitJSON
	// ParseFruitJSON (record []string) (*basepb.Fruit, error)
}

type fruitParser struct{}

func NewFruitParser() FruitParser {
	return &fruitParser{}
}
