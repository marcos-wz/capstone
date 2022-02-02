package controller

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/labstack/echo"
	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGetFruitsFilter(t *testing.T) {
	var testCases = []struct {
		name        string
		filterParam string
		valueParam  string
		respCode    int
		respBody    string
	}{
		{
			"Should return the fruits filtered, `Code: 200 - Fruits Filtered`",
			"id",
			"2",
			http.StatusOK,
			`{"Code":400,"Message":`,
		},
	}
	// NOTE: should i use echo or stdlib library ??
	// Otherwise, this will be a layer depended of the ECHO http framework,
	// But at the end of the day, this project use the ECHO http framework
	e := echo.New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ************
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/fruit/:filter/:value")
			c.SetParamNames("filter", "value")
			c.SetParamValues(tc.filterParam, tc.valueParam)
			// ************
			// csvfile := "../../data/fruits-test-ok.csv"
			// csvfile := "../../data/fruits-test-ok1.csv"
			csvfile := "../../data/fruits-test-error.csv"
			repo := repository.NewReaderRepo(csvfile)
			svc := service.NewFilterService(repo)
			h := NewFilterController(svc)

			// Assertions
			if assert.NoError(t, h.GetFruitsFilter(c)) {
				t.Log("Resp CODE:", rec.Code)
				t.Logf("Resp BODY: %s", rec.Body)
				// assert.Equal(t, http.StatusOK, rec.Code)
				// assert.Equal(t, userJSON, rec.Body.String())
			}
		})
	}

}
