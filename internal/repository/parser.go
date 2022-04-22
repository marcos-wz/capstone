package repository

import (
	"fmt"
	"github.com/go-playground/validator"
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
//		- input validation parsed array of errors
func parseFruitCSV(record []string, numFields int) (*basepb.Fruit, []error) {
	log.Println("REPO: parse fruit to csv starting...")
	// Values initialization
	fruit := &basepb.Fruit{}
	values := make([]string, numFields)
	copy(values, record)
	errors := make([]error, numFields)

	// Fruit Instance parse and validator ---------------------------------------------
	// use a single instance of Validate, it caches struct info
	var validate *validator.Validate
	// 0 - ID
	if err := validate.Var(values[0], "required,numeric,gt=0"); err != nil {
		log.Printf("REPO ERROR: CSV parser: ID %q: %v", values[0], err)
		errors[0] = err
	} else {
		ID, _ := strconv.ParseUint(values[0], 10, 32)
		fruit.Id = uint32(ID)
	}
	// 1 - NAME
	if err := validate.Var(values[1], "required,alpha,gt=2"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: NAME %q: %v", values[1], err)
		errors[1] = err
	} else {
		fruit.Name = values[1]
	}
	// 2 - DESCRIPTION
	if err := validate.Var(values[2], "omitempty,printascii,gt=2"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: DESCRIPTION %q: %v", values[2], err)
		errors[2] = err
	} else {
		fruit.Description = values[2]
	}
	// 3 - COLOR
	if err := validate.Var(values[3], "required,alpha,gt=2"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: COLOR %q: %v", values[3], err)
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
	if err := validate.Var(values[5], "required,numeric"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: price %q: %v", values[5], err)
		errors[5] = err
	} else {
		price, _ := strconv.ParseFloat(values[5], 32)
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
	if err := validate.Var(values[7], "required,numeric"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: STOCK %q: %v", values[7], err)
		errors[7] = err
	} else {
		stock, _ := strconv.ParseUint(values[7], 10, 32)
		fruit.Stock = uint32(stock)
	}
	// 8 - CADUCATE DAYS
	if err := validate.Var(values[8], "required,numeric"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: CADUCATE DAYS %q: %v", values[8], err)
		errors[8] = err
	} else {
		caducateDays, _ := strconv.ParseUint(values[8], 10, 32)
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
	if err := validate.Var(values[10], "required,numeric"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: CREATE TIME %q: %v", values[10], err)
		errors[10] = err
	} else {
		fruit.CreateTime, _ = strconv.ParseUint(values[10], 10, 64)
	}
	// 11 - UPDATE TIME
	if err := validate.Var(values[11], "required,numeric"); err != nil {
		log.Printf("REPO-ERROR: CSV parser: UPDATE TIME %q: %v", values[11], err)
		errors[10] = err
	} else {
		fruit.UpdateTime, _ = strconv.ParseUint(values[11], 10, 64)
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
