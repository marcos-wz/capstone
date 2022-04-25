package parser

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_ParseCountry(t *testing.T) {
	var testCases = []struct {
		name     string
		country  string
		response basepb.Country
	}{
		{
			"Should return `MEXICO` country",
			"MEXICO",
			basepb.Country_COUNTRY_MEXICO,
		},
		{
			"Should return `BRAZIL` country",
			"BRAZIL",
			basepb.Country_COUNTRY_BRAZIL,
		},
		{
			"Should return `CANADA` country",
			"CANADA",
			basepb.Country_COUNTRY_CANADA,
		},
		{
			"Should return `USA` country",
			"USA",
			basepb.Country_COUNTRY_USA,
		},
		{
			"Should return `UNDEFINED` country",
			"fake-country",
			basepb.Country_COUNTRY_UNDEFINED,
		},
	}
	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := NewFruitParser().ParseCountry(tc.country)
			assert.Equal(t, tc.response, resp)
			t.Logf("Parsed country %q : %v", tc.country, resp)
		})
	}
}
