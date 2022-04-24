package repository

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/marcos-wz/capstone/proto/basepb"
	"log"
	"strconv"
	"strings"
)

// parseFruitCSV guarantee csv data integrity. It parses from csv records to fruit instance.
// Record parameters is a string array from csv file. Returns a parsed fruit instance or error.
func parseFruitCSV(record []string) (*basepb.Fruit, error) {
	if Debug {
		log.Println("REPO: parse fruit to csv starting...")
		log.Printf("REPO: parser record(%d): %v", len(record), record)
	}
	// Load record to fruit by index
	fruitRecord := &entity.FruitCSVRecord{}
	for index, value := range record {
		// fix out of range index error
		switch index {
		case 0:
			fruitRecord.Id = value
		case 1:
			fruitRecord.Name = value
		case 2:
			fruitRecord.Description = value
		case 3:
			fruitRecord.Color = value
		case 4:
			fruitRecord.Unit = strings.ToUpper(value)
		case 5:
			fruitRecord.Price = value
		case 6:
			fruitRecord.Currency = strings.ToUpper(value)
		case 7:
			fruitRecord.Stock = value
		case 8:
			fruitRecord.CaducateDays = value
		case 9:
			fruitRecord.Country = strings.ToUpper(value)
		case 10:
			fruitRecord.CreateTime = value
		case 11:
			fruitRecord.UpdateTime = value
		default:
			err := fmt.Errorf("record index undefined %q : %q", index, value)
			log.Printf("REPO-ERROR: %v", err)
			return nil, err
		}
	}
	if Debug {
		log.Printf("REPO: fruit record : %+v", fruitRecord)
	}
	// Input Validation
	if err := validator.New().Struct(fruitRecord); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		log.Printf("REPO-ERROR: Record ID(%v) parser validation: ", fruitRecord.Id)
		for _, e := range validationErrors {
			log.Printf("Field: %v, Value: %v, Tag: %v, Param: %v", e.StructField(), e.Value(), e.Tag(), e.Param())
		}
		return nil, err
	}

	// Loading Protobuf Fruit
	id, _ := strconv.ParseUint(fruitRecord.Id, 10, 32)
	price, _ := strconv.ParseFloat(fruitRecord.Price, 32)
	stock, _ := strconv.ParseUint(fruitRecord.Stock, 10, 32)
	caducateDays, _ := strconv.ParseUint(fruitRecord.CaducateDays, 10, 32)
	createTime, _ := strconv.ParseUint(fruitRecord.CreateTime, 10, 64)
	updateTime, _ := strconv.ParseUint(fruitRecord.UpdateTime, 10, 64)

	return &basepb.Fruit{
		Id:           uint32(id),
		Name:         fruitRecord.Name,
		Description:  fruitRecord.Description,
		Color:        fruitRecord.Color,
		Unit:         parseUnit(fruitRecord.Unit),
		Price:        float32(price),
		Currency:     parseCurrency(fruitRecord.Currency),
		Stock:        uint32(stock),
		CaducateDays: uint32(caducateDays),
		Country:      parseCountry(fruitRecord.Country),
		CreateTime:   createTime,
		UpdateTime:   updateTime,
	}, nil
}

// parseUnit returns a parsed protobuf unit
func parseUnit(unit string) basepb.Unit {
	switch unit {
	case "KG":
		return basepb.Unit_UNIT_KG
	case "LB":
		return basepb.Unit_UNIT_LB
	default:
		log.Printf("REPO-WARNING: parser unit: %q undefined", unit)
		return basepb.Unit_UNIT_UNDEFINED
	}
}

// parseCountry parse a country and returns a unit prototype
func parseCountry(country string) basepb.Country {
	switch country {
	case "MEXICO":
		return basepb.Country_COUNTRY_MEXICO
	case "BRAZIL":
		return basepb.Country_COUNTRY_BRAZIL
	case "CANADA":
		return basepb.Country_COUNTRY_CANADA
	case "USA":
		return basepb.Country_COUNTRY_USA
	default:
		log.Printf("REPO-WARNING: country parser: country %q undefined", country)
		return basepb.Country_COUNTRY_UNDEFINED
	}
}

// parseCurrency parse a currency and returns a currency prototype
func parseCurrency(currency string) basepb.Currency {
	switch currency {
	case "MXN":
		return basepb.Currency_CURRENCY_MXN
	case "BRL":
		return basepb.Currency_CURRENCY_BRL
	case "CAD":
		return basepb.Currency_CURRENCY_CAD
	case "USD":
		return basepb.Currency_CURRENCY_USD
	default:
		log.Printf("REPO-WARNING: currency parser : currency %q undefined", currency)
		return basepb.Currency_CURRENCY_UNDEFINED
	}
}
