package service

import (
	"errors"
	"github.com/marcos-wz/capstone/internal/service/mocks"
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/stretchr/testify/assert"
	"testing"
)

// FilterFruits - UNIT TEST ****************************

func TestFilter_FilterFruits(t *testing.T) {

	// Test data, just filters values
	var mockFruitsData = []*basepb.Fruit{
		{Id: 1, Name: "Naranja", Country: basepb.Country_COUNTRY_CANADA, Color: "Yellow"},
		{Id: 2, Name: "manzana", Country: basepb.Country_COUNTRY_MEXICO, Color: "Red"},
		{Id: 3, Name: "Pera", Country: basepb.Country_COUNTRY_USA, Color: "Green"},
		{Id: 4, Name: "Platano", Country: basepb.Country_COUNTRY_USA, Color: "yellow"},
	}
	// test use cases
	var testCases = []struct {
		name         string
		filter       *filterpb.FilterRequest
		repoResponse []*basepb.Fruit
		repoErr      error
		response     []*basepb.Fruit
		err          error
	}{
		{
			"Should return the fruit filtered by ID 1, no errors",
			&filterpb.FilterRequest{
				Filter: filterpb.FiltersAllowed_FILTER_ID,
				Value:  "1",
			},
			mockFruitsData,
			nil,
			[]*basepb.Fruit{
				{Id: 1, Name: "Naranja", Country: basepb.Country_COUNTRY_CANADA, Color: "Yellow"},
			},
			nil,
		},
		{
			"Should return the fruit filtered by NAME, no errors",
			&filterpb.FilterRequest{Filter: filterpb.FiltersAllowed_FILTER_NAME, Value: "pera"},
			mockFruitsData,
			nil,
			[]*basepb.Fruit{
				{Id: 3, Name: "Pera", Country: basepb.Country_COUNTRY_USA, Color: "Green"},
			},
			nil,
		},
		{
			"Should return the fruit filtered by COLOR, no errors",
			&filterpb.FilterRequest{Filter: filterpb.FiltersAllowed_FILTER_COLOR, Value: "yellow"},
			mockFruitsData,
			nil,
			[]*basepb.Fruit{
				{Id: 1, Name: "Naranja", Country: basepb.Country_COUNTRY_CANADA, Color: "Yellow"},
				{Id: 4, Name: "Platano", Country: basepb.Country_COUNTRY_USA, Color: "yellow"},
			},
			nil,
		},
		{
			"Should return the fruit filtered by COUNTRY, no errors",
			&filterpb.FilterRequest{Filter: filterpb.FiltersAllowed_FILTER_COUNTRY, Value: "usa"},
			mockFruitsData,
			nil,
			[]*basepb.Fruit{
				{Id: 3, Name: "Pera", Country: basepb.Country_COUNTRY_USA, Color: "Green"},
				{Id: 4, Name: "Platano", Country: basepb.Country_COUNTRY_USA, Color: "yellow"},
			},
			nil,
		},
		// Repository error
		{
			"Should return repository error",
			&filterpb.FilterRequest{Filter: filterpb.FiltersAllowed_FILTER_ID, Value: "1"},
			mockFruitsData,
			errors.New("fake repo error"),
			nil,
			errors.New("fake repo error"),
		},
	}

	// RUN TESTS --------------------------------------
	//Debug = true
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// MOCK
			mock := mocks.NewFruitRepo(t)
			mock.On("ReadFruits").Return(tc.repoResponse, tc.repoErr)
			//// SERVICE
			service := NewFruitService(mock)
			fruits, err := service.FilterFruits(tc.filter)
			if Debug {
				t.Log("Total fruits:", len(fruits))
				t.Logf("Err: %v", err)
				for _, f := range fruits {
					t.Logf("%+v", f)
				}
			}
			assert.Equal(t, tc.err, err)
			assert.Len(t, tc.response, len(fruits))
			for i := 0; i < len(fruits); i++ {
				assert.EqualValues(t, tc.response[i], fruits[i])
			}
		})

	}
}
