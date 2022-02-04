package service

import (
	"fmt"
	"testing"

	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock
type mockReaderRepo struct {
	mock.Mock
}

func (mr *mockReaderRepo) ReadFruits() ([]entity.Fruit, error) {
	arg := mr.Called()
	return arg.Get(0).([]entity.Fruit), arg.Error(1)
}

// Test data, just filters values
var mockFruitsData = []entity.Fruit{
	{ID: 1, Name: "Naranja", Country: "Canada", Color: "Yellow"},
	{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"},
	{ID: 3, Name: "Pera", Country: "usa", Color: "Green"},
	{ID: 4, Name: "Platano", Country: "USA", Color: "yellow"},
}

// UNIT TEST ****************************

func TestFilter_FilterFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filter   *entity.FruitsFilterParams
		repoResp []entity.Fruit
		repoErr  error
		response []entity.Fruit
		err      error
	}{
		{
			"Should return the fruit filtered by ID, no errors",
			&entity.FruitsFilterParams{Filter: "id", Value: "1"},
			mockFruitsData,
			nil,
			[]entity.Fruit{
				{ID: 1, Name: "Naranja", Country: "Canada", Color: "Yellow"},
			},
			nil,
		},
		{
			"Should return error, Invalid Filter ID: invalid syntax",
			&entity.FruitsFilterParams{Filter: "id", Value: "badvalue"},
			mockFruitsData,
			nil,
			nil,
			fmt.Errorf("invalid ID filter(badvalue): strconv.Atoi: parsing \"badvalue\": invalid syntax"),
		},
		{
			"Should return the fruit filtered by NAME, no errors",
			&entity.FruitsFilterParams{Filter: "name", Value: "pera"},
			mockFruitsData,
			nil,
			[]entity.Fruit{{ID: 3, Name: "Pera", Country: "usa", Color: "Green"}},
			nil,
		},
		{
			"Should return the fruit filtered by COLOR, no errors",
			&entity.FruitsFilterParams{Filter: "color", Value: "yellow"},
			mockFruitsData,
			nil,
			[]entity.Fruit{
				{ID: 1, Name: "Naranja", Country: "Canada", Color: "Yellow"},
				{ID: 4, Name: "Platano", Country: "USA", Color: "yellow"},
			},
			nil,
		},
		{
			"Should return the fruit filtered by COUNTRY, no errors",
			&entity.FruitsFilterParams{Filter: "country", Value: "USA"},
			mockFruitsData,
			nil,
			[]entity.Fruit{
				{ID: 3, Name: "Pera", Country: "usa", Color: "Green"},
				{ID: 4, Name: "Platano", Country: "USA", Color: "yellow"},
			},
			nil,
		},
		{
			"Should return all fruits, no errors",
			&entity.FruitsFilterParams{Filter: "all", Value: ""},
			mockFruitsData,
			nil,
			[]entity.Fruit{
				{ID: 1, Name: "Naranja", Country: "Canada", Color: "Yellow"},
				{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"},
				{ID: 3, Name: "Pera", Country: "usa", Color: "Green"},
				{ID: 4, Name: "Platano", Country: "USA", Color: "yellow"},
			},
			nil,
		},
		{
			"Should return error, Invalid Filter: undefined filter",
			&entity.FruitsFilterParams{Filter: "badfilter", Value: "badvalue"},
			mockFruitsData,
			nil,
			nil,
			fmt.Errorf("undefined filter(badfilter): badvalue"),
		},
		// Repository error: no such file
		{
			"Should return repository error, open : no such file or directory",
			&entity.FruitsFilterParams{Filter: "id", Value: "1"},
			nil,
			fmt.Errorf("open : no such file or directory"),
			nil,
			fmt.Errorf("open : no such file or directory"),
		},
		// Repository error: parser error
		{
			"Should return repository error, parser error:",
			&entity.FruitsFilterParams{Filter: "id", Value: "1"},
			mockFruitsData,
			fmt.Errorf("parser error: "),
			[]entity.Fruit{
				{ID: 1, Name: "Naranja", Country: "Canada", Color: "Yellow"},
			},
			fmt.Errorf("parser error: "),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// MOCK
			mock := mockReaderRepo{}
			mock.On("ReadFruits").Return(tc.repoResp, tc.repoErr)
			// SERVICE
			service := NewFilterService(&mock)
			fruits, err := service.GetFilteredFruits(tc.filter)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.response, fruits)
			if len(fruits) > 0 {
				t.Log("Total fruits:", len(fruits))
				for _, f := range fruits {
					t.Logf("%+v", f)
				}
			}
		})

	}
}
