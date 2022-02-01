package service

import (
	"fmt"
	"testing"

	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MOCK *************************************

type mockFilter struct {
	mock.Mock
}

func (mr *mockFilter) ReadFruits() ([]entity.Fruit, error) {
	arg := mr.Called()
	return arg.Get(0).([]entity.Fruit), arg.Error(1)
}

// TEST DATA
var mockFruitsData = []entity.Fruit{
	{
		ID:      1,
		Name:    "Naranja",
		Country: "Canada",
		Color:   "Yellow",
	},
	{
		ID:      2,
		Name:    "Manzana",
		Country: "Mexico",
		Color:   "Red",
	},
	{
		ID:      3,
		Name:    "Pera",
		Country: "USA",
		Color:   "Green",
	},
	{
		ID:      4,
		Name:    "Platano",
		Country: "USA",
		Color:   "Yellow",
	},
}

// UNIT TEST ****************************

func TestFilterFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filter   string
		value    string
		repoResp []entity.Fruit
		repoErr  error
		response []entity.Fruit
		err      error
	}{
		{
			"Should return the fruit filtered by ID, no errors",
			"id",
			"1",
			mockFruitsData,
			nil,
			[]entity.Fruit{mockFruitsData[0]},
			nil,
		},
		{
			"Should return error, Invalid Filter: id - badvalue",
			"id",
			"badvalue",
			mockFruitsData,
			nil,
			nil,
			fmt.Errorf("invalid ID filter: badvalue - strconv.Atoi: parsing \"badvalue\": invalid syntax"),
		},
		{
			"Should return the fruit filtered by NAME, no errors",
			"name",
			"pera",
			mockFruitsData,
			nil,
			[]entity.Fruit{mockFruitsData[2]},
			nil,
		},
		{
			"Should return the fruit filtered by COLOR, no errors",
			"color",
			"yellow",
			mockFruitsData,
			nil,
			[]entity.Fruit{mockFruitsData[0], mockFruitsData[3]},
			nil,
		},
		{
			"Should return the fruit filtered by COUNTRY, no errors",
			"country",
			"USA",
			mockFruitsData,
			nil,
			[]entity.Fruit{mockFruitsData[2], mockFruitsData[3]},
			nil,
		},
		{
			"Should return all fruits, no errors",
			"all",
			"",
			mockFruitsData,
			nil,
			mockFruitsData,
			nil,
		},
		{
			"Should return error, Invalid Filter: badfilter - badvalue",
			"badfilter",
			"badvalue",
			mockFruitsData,
			nil,
			nil,
			fmt.Errorf("invalid filter: badfilter - badvalue"),
		},
		// REPOSITORY ERRORS
		{
			"Should return repository error, open : no such file or directory",
			"id",
			"1",
			nil,
			fmt.Errorf("open : no such file or directory"),
			nil,
			fmt.Errorf("open : no such file or directory"),
		},
		// REPOSITORY ERRORS
		{
			"Should return repository error, parser error:",
			"id",
			"1",
			mockFruitsData,
			fmt.Errorf("parser error:"),
			[]entity.Fruit{mockFruitsData[0]},
			fmt.Errorf("parser error:"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// MOCK
			mock := mockFilter{}
			mock.On("ReadFruits").Return(tc.repoResp, tc.repoErr)
			// SERVICE
			service := NewFilterService(&mock)
			fruits, err := service.FilterFruits(tc.filter, tc.value)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.response, fruits)
		})

	}
}
