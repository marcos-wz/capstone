package repository

import (
	"fmt"
	"github.com/marcos-wz/capstone/proto/basepb"
	"log"
	"strconv"
	"strings"
)

// parseFruitCSV guarantee csv data integrity. It parses from csv records to fruit instance
// This is an input data method. Always returns a fruit instance.
// If an error occurs, the default type value is set.
// PARAMETERS:
// 		- record: string array from csv file,
// 		- numFields: The number of fruit fields
// RETURNS:
//		- parsed fruit instance
//		- validation parsed array of errors
func parseFruitCSV(record []string, numFields int) (*basepb.Fruit, []error) {
	log.Println("REPO: parse fruit to csv starting...")
	// Values initialization
	fruit := &basepb.Fruit{}
	values := make([]string, numFields)
	copy(values, record)
	errors := make([]error, numFields)

	// Fruit Instance ---------------------------------------------
	// 0 - ID
	ID, err := strconv.ParseUint(values[0], 10, 32)
	if err != nil {
		log.Printf("ERROR-Parser: parsing ID %q: %v", ID, err)
		errors[0] = err
	} else {
		fruit.Id = uint32(ID)
	}
	// 1 - NAME
	if values[1] == "" {
		err := fmt.Errorf("empty string")
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[1] = err
	} else {
		fruit.Name = values[1]
	}
	// 2 - DESCRIPTION
	if values[2] == "" {
		err := fmt.Errorf("empty string")
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[2] = err
	} else {
		fruit.Description = values[2]
	}
	// 3 - COLOR
	if values[3] == "" {
		err := fmt.Errorf("empty string")
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[3] = err
	} else {
		fruit.Color = values[3]
	}
	// 4 - UNIT
	fruit.Unit = parseUnit(values[4])
	if fruit.Unit == basepb.Unit_UNIT_UNDEFINED {
		err := fmt.Errorf("unit %q is undefined", values[4])
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[4] = err
	}
	// 5 - PRICE
	price, err := strconv.ParseFloat(values[5], 32)
	if err != nil {
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[5] = err
	} else {
		fruit.Price = float32(price)
	}
	// 6 - CURRENCY
	fruit.Currency = parseCurrency(values[6])
	if fruit.Currency == basepb.Currency_CURRENCY_UNDEFINED {
		err := fmt.Errorf("concurrency %q is undefined", values[6])
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[6] = err
	}
	// 7 - STOCK
	stock, err := strconv.ParseUint(values[7], 10, 32)
	if err != nil {
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[7] = err
	} else {
		fruit.Stock = uint32(stock)
	}
	// 8 - CADUCATE DAYS
	caducateDays, err := strconv.ParseUint(values[8], 10, 32)
	if err != nil {
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[8] = err
	} else {
		fruit.CaducateDays = uint32(caducateDays)
	}
	// 9 - COUNTRY
	fruit.Country = parseCountry(values[9])
	if fruit.Country == basepb.Country_COUNTRY_UNDEFINED {
		err := fmt.Errorf("country %q is undefined", values[9])
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[8] = err
	}
	// 10 - CREATE TIME
	fruit.CreateTime, err = strconv.ParseUint(values[10], 10, 64)
	if err != nil {
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[10] = err
	}
	// 11 - UPDATE TIME
	fruit.UpdateTime, err = strconv.ParseUint(values[11], 10, 64)
	if err != nil {
		log.Printf("REPO-ERROR: CSV parser: %v", err)
		errors[10] = err
	}

	return fruit, errors
}

// parseUnit parse a unit, and returns a unit prototype
func parseUnit(unit string) basepb.Unit {
	switch strings.ToUpper(unit) {
	case "KG":
		return basepb.Unit_UNIT_KG
	case "LB":
		return basepb.Unit_UNIT_LB
	default:
		log.Printf("REPO-ERROR: units parser: unit %v undefined", unit)
		return basepb.Unit_UNIT_UNDEFINED
	}
}

// parseCountry parse a country and returns a unit prototype
func parseCountry(country string) basepb.Country {
	switch strings.ToUpper(country) {
	case "MEXICO":
		return basepb.Country_COUNTRY_MEXICO
	case "BRAZIL":
		return basepb.Country_COUNTRY_BRAZIL
	case "CANADA":
		return basepb.Country_COUNTRY_CANADA
	case "USA":
		return basepb.Country_COUNTRY_USA
	default:
		return basepb.Country_COUNTRY_UNDEFINED
	}
}

// parseCurrency parse a currency and returns a currency prototype
func parseCurrency(currency string) basepb.Currency {
	switch strings.ToUpper(currency) {
	case "MXN":
		return basepb.Currency_CURRENCY_MXN
	case "BRL":
		return basepb.Currency_CURRENCY_BRL
	case "CAD":
		return basepb.Currency_CURRENCY_CAD
	case "USD":
		return basepb.Currency_CURRENCY_USD
	default:
		return basepb.Currency_CURRENCY_UNDEFINED
	}
}
