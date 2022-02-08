package controller

import (
	"encoding/json"
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

func (mr *mockFilterService) GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, *entity.FruitsFilterError) {
	arg := mr.Called()
	return arg.Get(0).([]entity.Fruit), arg.Get(1).(*entity.FruitsFilterError)
}

// UNIT TEST ***************************************

func TestFilter_FilterFruit(t *testing.T) {

	e := echo.New() // Echo http server instance
	// Test use cases
	var testCases = []struct {
		name            string
		params          *entity.FruitsFilterParams
		serviceResponse []entity.Fruit
		serviceError    *entity.FruitsFilterError
		responseCode    int
		responseBody    string
	}{
		{
			"Should return the fruit ID 2 with Code: 200",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			[]entity.Fruit{
				{ID: 2, Name: "Manzana", Country: "Mexico", Color: "Red"},
			},
			nil,
			http.StatusOK,
			"{\"fruits\":[{\"id\":2,\"name\":\"Manzana\",\"description\":\"\",\"color\":\"Red\",\"unit\":\"\",\"price\":0,\"stock\":0,\"caducate_days\":0,\"country\":\"Mexico\",\"created_at\":\"0001-01-01T00:00:00Z\"}]}\n",
		},
		{
			"Should return error empty filter, Error:Field validation with Code: 422",
			&entity.FruitsFilterParams{},
			nil,
			nil,
			http.StatusUnprocessableEntity,
			"{\"message\":\"Key: 'FruitsFilterParams.Filter' Error:Field validation for 'Filter' failed on the 'oneof' tag\\nKey: 'FruitsFilterParams.Value' Error:Field validation for 'Value' failed on the 'required' tag\"}\n",
		},
		{
			"Should return error filter value, Error:Field validation with Code: 422",
			&entity.FruitsFilterParams{Filter: "id"},
			nil,
			nil,
			http.StatusUnprocessableEntity,
			"{\"message\":\"Key: 'FruitsFilterParams.Value' Error:Field validation for 'Value' failed on the 'required' tag\"}\n",
		},
		{
			"Should return filter validation error, Error:Field validation with Code: 422",
			&entity.FruitsFilterParams{Filter: "badfilter", Value: "12"},
			nil,
			nil,
			http.StatusUnprocessableEntity,
			"{\"message\":\"Key: 'FruitsFilterParams.Filter' Error:Field validation for 'Filter' failed on the 'oneof' tag\"}\n",
		},
		// REPOSITORY USE CASE ERRORS
		{
			"Should return repository internal server error: Code 500, no such file or directory",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			nil,
			&entity.FruitsFilterError{
				Type:  "Repo.FileError",
				Error: errors.New("no such file or directory"),
			},
			http.StatusInternalServerError,
			"{\"message\":\"no such file or directory\"}\n",
		},
		// {
		// 	"Should return repository parser error: Code 206, partial content",
		// 	&entity.FruitsFilterParams{Filter: "id", Value: "2"},
		// 	[]entity.Fruit{
		// 		{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"},
		// 	},
		// 	errors.New("parser error: "),
		// 	http.StatusPartialContent,
		// 	"{\"fruits\":[{\"id\":2,\"name\":\"manzana\",\"description\":\"\",\"color\":\"Red\",\"unit\":\"\",\"price\":0,\"stock\":0,\"caducate\":0,\"country\":\"Mexico\",\"CreatedAt\":\"0001-01-01T00:00:00Z\"}],\"parser_error\":\"parser error: \"}\n",
		// },
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
			// JSON RESPONSE
			bodyJson := make(map[string]interface{}) // unmarshall json reponses intance

			// Assertions
			if assert.NoError(t, ctrl.FilterFruit(c)) {

				assert.Equal(t, tc.responseCode, rec.Code)
				assert.Equal(t, tc.responseBody, rec.Body.String())
				assert.NoError(t, json.NewDecoder(rec.Body).Decode(&bodyJson))
				// Response validations
				switch rec.Code {
				case http.StatusOK:
					fruits, ok := bodyJson["fruits"]
					if assert.True(t, ok) {
						assert.NotEmpty(t, fruits)
						assert.Len(t, fruits, len(tc.serviceResponse))
					}
				case http.StatusUnprocessableEntity:
					message, ok := bodyJson["message"]
					if assert.True(t, ok) {
						assert.NotEmpty(t, message)
					}
				case http.StatusInternalServerError:
					message, ok := bodyJson["message"]
					if assert.True(t, ok) {
						assert.NotEmpty(t, message)
					}
				case http.StatusPartialContent:
					t.Log("Unprocessable Entity Val")
				case http.StatusBadRequest:
					t.Log("Unprocessable Entity Val")
				default:
					t.Errorf("The status code %d, should be defined: %s", rec.Code, rec.Body)
				}
				t.Log("Resp CODE:", rec.Code)
				t.Logf("Resp BODY Json: %v", bodyJson)
				// t.Logf("Resp BODY String: %v", rec.Body.String())
			}
		})
	}

}
