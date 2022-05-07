package repository

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"log"
	"strconv"
	"strings"
)

func ParseCountry(country string) basepb.Country {
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
		log.Printf("PARSER-WARNING: country parser: country %q undefined", country)
		return basepb.Country_COUNTRY_UNDEFINED
	}
}

func ParseCurrency(currency string) basepb.Currency {
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
		log.Printf("PARSER-WARNING: currency parser : currency %q undefined", currency)
		return basepb.Currency_CURRENCY_UNDEFINED
	}
}

func ParseFilter(filter string) filterpb.FiltersAllowed {
	switch strings.ToUpper(filter) {
	case "ID":
		return filterpb.FiltersAllowed_FILTER_ID
	case "NAME":
		return filterpb.FiltersAllowed_FILTER_NAME
	case "COLOR":
		return filterpb.FiltersAllowed_FILTER_COLOR
	case "COUNTRY":
		return filterpb.FiltersAllowed_FILTER_COUNTRY
	default:
		log.Printf("PARSER-WARNING: filter parser : filter %q undefined", filter)
		return filterpb.FiltersAllowed_FILTER_UNDEFINED
	}
}

func ParseFruitRecord(record []string) (*basepb.Fruit, error) {
	if DebugLevel >= 2 {
		log.Println("REPO: parse fruit to csv starting...")
		log.Printf("REPO: parser record(%d): %v", len(record), record)
	}
	// Load record to fruit by field index
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
			errIndex := fmt.Errorf("%q : %q", index, value)
			log.Printf("PARSER-ERROR: %v", errIndex)
			return nil, &entity.FruitError{
				Type: ErrParseFruitCSVIndex, Desc: ErrDesc[ErrParseFruitCSVIndex], Err: errIndex,
			}
		}
	}
	if DebugLevel >= 2 {
		log.Printf("REPO: fruit record : %+v", fruitRecord)
	}
	// Record validator validations
	if err := validator.New().Struct(fruitRecord); err != nil {
		log.Printf("PARSER-ERROR: %v", err)
		if errValidation, ok := err.(validator.ValidationErrors); ok {
			return nil, &entity.FruitError{
				Type: ErrParseFruitCSVValidation,
				Desc: ErrDesc[ErrParseFruitCSVValidation], Err: errValidation,
			}
		}
		return nil, err
	}

	// Mapping values from csv record to proto Fruit
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
		Unit:         ParseUnit(fruitRecord.Unit),
		Price:        float32(price),
		Currency:     ParseCurrency(fruitRecord.Currency),
		Stock:        uint32(stock),
		CaducateDays: uint32(caducateDays),
		Country:      ParseCountry(fruitRecord.Country),
		CreateTime:   createTime,
		UpdateTime:   updateTime,
	}, nil
}

func ParseUnit(unit string) basepb.Unit {
	switch unit {
	case "KG":
		return basepb.Unit_UNIT_KG
	case "LB":
		return basepb.Unit_UNIT_LB
	default:
		log.Printf("PARSER-WARNING: parser unit: %q undefined", unit)
		return basepb.Unit_UNIT_UNDEFINED
	}
}
