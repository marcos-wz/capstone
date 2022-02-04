package controller

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/labstack/echo"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MOCK *******************************************

type mockFilterService struct {
	mock.Mock
}

func (mr *mockFilterService) GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, error) {
	arg := mr.Called()
	return arg.Get(0).([]entity.Fruit), arg.Error(1)
}

var mockFruitsData = []entity.Fruit{
	{ID: 1, Name: "Naranja", Country: "Canada", Color: "Yellow"},
	{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"},
	{ID: 3, Name: "Pera", Country: "usa", Color: "Green"},
	{ID: 4, Name: "Platano", Country: "USA", Color: "yellow"},
}

// UNIT TEST ***************************************

func TestFilter_FilterFruit(t *testing.T) {
	// Echo http server instance
	e := echo.New()
	// Test use cases
	var testCases = []struct {
		name            string
		params          *entity.FruitsFilterParams
		serviceResponse []entity.Fruit
		serviceError    error
		responseCode    int
		responseBody    string
	}{
		{
			"Should return the fruits filtered, `Code: 200 - Fruits Filtered`",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			[]entity.Fruit{
				{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"},
			},
			nil,
			http.StatusOK,
			`{"Code":400,"Message":`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Request
			url := "/v1/fruit/:filter/:value"
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("filter", "value")
			c.SetParamValues(tc.params.Filter, tc.params.Value)
			// Mock
			mock := mockFilterService{}
			mock.On("GetFilteredFruits").Return(tc.serviceResponse, tc.serviceError)
			// Controller
			ctrl := NewFilterController(&mock)

			// Assertions
			if assert.NoError(t, ctrl.FilterFruit(c)) {
				t.Log("Resp CODE:", rec.Code)
				t.Logf("Resp BODY: %s", rec.Body)
				assert.Equal(t, tc.responseCode, rec.Code)
				// assert.Equal(t, userJSON, rec.Body.String())
			}
		})
	}

}
