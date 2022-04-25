package parser

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/marcos-wz/capstone/proto/basepb"
	"log"
	"strconv"
	"strings"
)

func (fp *fruitParser) ParseFruitCSVRecord(record []string) (*basepb.Fruit, error) {
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
		Unit:         fp.ParseUnit(fruitRecord.Unit),
		Price:        float32(price),
		Currency:     fp.ParseCurrency(fruitRecord.Currency),
		Stock:        uint32(stock),
		CaducateDays: uint32(caducateDays),
		Country:      fp.ParseCountry(fruitRecord.Country),
		CreateTime:   createTime,
		UpdateTime:   updateTime,
	}, nil
}
