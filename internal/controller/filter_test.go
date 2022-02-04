package controller

import (
	"errors"
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
			"Should return the fruit ID 2 with Code: 200",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			[]entity.Fruit{
				{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"},
			},
			nil,
			http.StatusOK,
			"{\"fruits\":[{\"id\":2,\"name\":\"manzana\",\"description\":\"\",\"color\":\"Red\",\"unit\":\"\",\"price\":0,\"stock\":0,\"caducate\":0,\"country\":\"Mexico\",\"CreatedAt\":\"0001-01-01T00:00:00Z\"}],\"parser_error\":\"\"}\n",
		},
		{
			"Should return error empty filter, Error:Field validation with Code: 422",
			&entity.FruitsFilterParams{},
			nil,
			nil,
			http.StatusUnprocessableEntity,
			"\"Key: 'FruitsFilterParams.Filter' Error:Field validation for 'Filter' failed on the 'oneof' tag\\nKey: 'FruitsFilterParams.Value' Error:Field validation for 'Value' failed on the 'alphanum' tag\"\n",
		},
		{
			"Should return error filter value, Error:Field validation with Code: 422",
			&entity.FruitsFilterParams{Filter: "id"},
			nil,
			nil,
			http.StatusUnprocessableEntity,
			"\"Key: 'FruitsFilterParams.Value' Error:Field validation for 'Value' failed on the 'alphanum' tag\"\n",
		},
		{
			"Should return repository internal server error: Code 500, no such file or directory",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			nil,
			errors.New("no such file or directory"),
			http.StatusInternalServerError,
			"\"no such file or directory\"\n",
		},
		{
			"Should return repository parser error: Code 206, partial content",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			[]entity.Fruit{
				{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"},
			},
			errors.New("parser error: "),
			http.StatusPartialContent,
			"{\"fruits\":[{\"id\":2,\"name\":\"manzana\",\"description\":\"\",\"color\":\"Red\",\"unit\":\"\",\"price\":0,\"stock\":0,\"caducate\":0,\"country\":\"Mexico\",\"CreatedAt\":\"0001-01-01T00:00:00Z\"}],\"parser_error\":\"parser error: \"}\n",
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
				assert.Equal(t, tc.responseCode, rec.Code)
				assert.Equal(t, tc.responseBody, rec.Body.String())
				t.Log("Resp CODE:", rec.Code)
				t.Logf("Resp BODY: %s", rec.Body)
			}
		})
	}

}
