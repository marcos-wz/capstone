package repository

import (
	"fmt"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_ReadFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		response []*basepb.Fruit
		err      string
	}{
		{
			"Should return all fruits: no errors",
			"../../data/test/csv/fruits-test-ok.csv",
			[]*basepb.Fruit{
				{Id: 1, Name: "Pera", Description: "Fruta Tropical", Color: "green", Unit: basepb.Unit_UNIT_KG, Currency: basepb.Currency_CURRENCY_MXN, Price: 5.50, Stock: 1, CaducateDays: 1, Country: basepb.Country_COUNTRY_MEXICO, CreateTime: 1642802058, UpdateTime: 1647899658},
				{Id: 2, Name: "Manzana", Description: "Fruta tropical", Color: "red", Unit: basepb.Unit_UNIT_KG, Currency: basepb.Currency_CURRENCY_BRL, Price: 2, Stock: 1, CaducateDays: 1, Country: basepb.Country_COUNTRY_BRAZIL, CreateTime: 1642802058, UpdateTime: 1647899658},
				{Id: 3, Name: "Platano", Description: "Fruta tropical", Color: "yellow", Unit: basepb.Unit_UNIT_LB, Currency: basepb.Currency_CURRENCY_USD, Price: 20, Stock: 1, CaducateDays: 1, Country: basepb.Country_COUNTRY_USA, CreateTime: 1642802058, UpdateTime: 1647899658},
				{Id: 4, Name: "Mandarina", Description: "Fruta tropical", Color: "orange", Unit: basepb.Unit_UNIT_LB, Currency: basepb.Currency_CURRENCY_CAD, Price: 20, Stock: 1, CaducateDays: 1, Country: basepb.Country_COUNTRY_CANADA, CreateTime: 1642802058, UpdateTime: 1647899658},
				{Id: 5, Name: "Naranja", Description: "Fruta tropical", Color: "yellow", Unit: basepb.Unit_UNIT_LB, Currency: basepb.Currency_CURRENCY_USD, Price: 20, Stock: 1, CaducateDays: 1, Country: basepb.Country_COUNTRY_USA, CreateTime: 1642802058, UpdateTime: 1647899658},
			},
			"<nil>",
		},
		{
			"Should return error `no such file or directory`",
			"",
			nil,
			"csv file error: open : no such file or directory",
		},
		{
			"Should return empty fruits list, with parser error",
			"../../data/test/csv/fruits-test-error.csv",
			nil,
			"parse fruit csv record validation error: Key: 'FruitCSVRecord.Id' Error:Field validation for 'Id' failed on the 'required' tag\nKey: 'FruitCSVRecord.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag",
		},
		{
			"Should return empty fruits list, with parser error",
			"../../data/test/csv/fruits-test-partial-parser-error.csv",
			nil,
			"parse fruit csv record validation error: Key: 'FruitCSVRecord.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag\nKey: 'FruitCSVRecord.Currency' Error:Field validation for 'Currency' failed on the 'oneof' tag",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewFruitRepo(tc.filePath, "")
			fruits, err := repo.ReadFruits()
			assert.Equal(t, tc.err, fmt.Sprintf("%v", err))
			// Response
			assert.Len(t, fruits, len(tc.response))
			assert.EqualValues(t, tc.response, fruits)
			//t.Logf("Fruit: %v", fruits)
			//t.Logf("Error: %v", err)
		})
	}
}

//
