package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/labstack/echo"
	"github.com/marcos-wz/capstone/internal/controller/mocks"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
)

// FILTER FRUIT - UNIT TEST ***************************************

func TestFilter_FilterFruit(t *testing.T) {
	// Echo http server instance
	e := echo.New()
	// Test use cases
	var testCases = []struct {
		name            string
		params          *entity.FruitsFilterParams
		serviceResponse []entity.Fruit
		serviceError    *entity.FruitFilterError
		responseCode    int
		responseBody    string
	}{
		{
			"Should return the fruit ID 2 with Code: 200",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			[]entity.Fruit{{ID: 2, Name: "Manzana", Country: "Mexico", Color: "Red"}},
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
			&entity.FruitFilterError{
				Type:  "Repo.FileError",
				Error: errors.New("no such file or directory"),
			},
			http.StatusInternalServerError,
			"{\"message\":\"no such file or directory\"}\n",
		},
		{
			"Should return repository parser error: Code 206, partial content",
			&entity.FruitsFilterParams{Filter: "id", Value: "2"},
			[]entity.Fruit{{ID: 2, Name: "manzana", Country: "Mexico", Color: "Red"}},
			&entity.FruitFilterError{
				Type:  "Repo.ParserError",
				Error: errors.New("reader repository, parse fruit errors found"),
				ParserErrors: []entity.ParseFruitRecordCSVError{
					{Record: 1, Errors: []entity.ParseFruitFieldCSVError{
						{Field: "ID", Value: "0", Validation: "required", Error: "Key: 'Fruit.ID' Error:Field validation for 'ID' failed on the 'required' tag"},
						{Field: "Unit", Value: "lbs", Validation: "oneof", Error: "Key: 'Fruit.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag"},
					}},
				},
			},
			http.StatusPartialContent,
			"{\"fruits\":[{\"id\":2,\"name\":\"manzana\",\"description\":\"\",\"color\":\"Red\",\"unit\":\"\",\"price\":0,\"stock\":0,\"caducate_days\":0,\"country\":\"Mexico\",\"created_at\":\"0001-01-01T00:00:00Z\"}],\"parser_errors\":[{\"record\":1,\"errors\":[{\"field\":\"ID\",\"value\":\"0\",\"validation\":\"required\",\"error\":\"Key: 'Fruit.ID' Error:Field validation for 'ID' failed on the 'required' tag\"},{\"field\":\"Unit\",\"value\":\"lbs\",\"validation\":\"oneof\",\"error\":\"Key: 'Fruit.Unit' Error:Field validation for 'Unit' failed on the 'oneof' tag\"}]}]}\n",
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
			mock := &mocks.FilterService{}
			mock.On("GetFilteredFruits", tc.params).Return(tc.serviceResponse, tc.serviceError)
			// Controller
			ctrl := NewFilterController(mock)

			// Assertions
			if assert.NoError(t, ctrl.FilterFruit(c)) {
				assert.Equal(t, tc.responseCode, rec.Code)
				assert.Equal(t, tc.responseBody, rec.Body.String())
				switch rec.Code {
				// Response OK : 200
				case http.StatusOK:
					resp := &entity.FruitFilterResponse{}
					assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), resp))
					assert.EqualValues(t, tc.serviceResponse, resp.Fruits)
					assert.Nil(t, resp.ParserErrors)
				// Response UNPROCESSABLE-ENTITY : 422
				case http.StatusUnprocessableEntity:
					resp, respTc := &entity.ErrorResponse{}, &entity.ErrorResponse{}
					assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), resp))
					assert.NoError(t, json.Unmarshal([]byte(tc.responseBody), respTc))
					assert.Equal(t, respTc, resp)
				// Response INTERNAL-SERVER : 500
				case http.StatusInternalServerError:
					resp := &entity.ErrorResponse{}
					assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), resp))
					assert.EqualError(t, tc.serviceError.Error, resp.Message)
				// Response PARTIAL-CONTENT : 206
				case http.StatusPartialContent:
					resp := &entity.FruitFilterResponse{}
					assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), resp))
					assert.EqualValues(t, tc.serviceResponse, resp.Fruits)
					assert.Equal(t, tc.serviceError.ParserErrors, resp.ParserErrors)
				default:
					t.Errorf("The status code %d, should be defined: %s", rec.Code, rec.Body)
				}
			}
		})
	}

}
