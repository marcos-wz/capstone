package parser

import (
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_ParseFilter(t *testing.T) {
	var testCases = []struct {
		name     string
		filter   string
		response filterpb.FiltersAllowed
	}{
		{
			"Should return `ID` filter",
			"id",
			filterpb.FiltersAllowed_FILTER_ID,
		},
		{
			"Should return `NAME` filter",
			"id",
			filterpb.FiltersAllowed_FILTER_ID,
		},
		{
			"Should return `COLOR` filter",
			"color",
			filterpb.FiltersAllowed_FILTER_COLOR,
		},
		{
			"Should return `COUNTRY` filter",
			"country",
			filterpb.FiltersAllowed_FILTER_COUNTRY,
		},
		{
			"Should return `UNDEFINED` filter",
			"fake-filter",
			filterpb.FiltersAllowed_FILTER_UNDEFINED,
		},
	}
	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := NewFruitParser().ParseFilter(tc.filter)
			assert.Equal(t, tc.response, resp)
			t.Logf("Parsed filter %q : %v", tc.filter, resp)
		})
	}
}
