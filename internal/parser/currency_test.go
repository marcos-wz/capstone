package parser

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_ParseCurrency(t *testing.T) {
	var testCases = []struct {
		name     string
		currency string
		response basepb.Currency
	}{
		{
			"Should return `MNX` currency",
			"MXN",
			basepb.Currency_CURRENCY_MXN,
		},
		{
			"Should return `BRL` currency",
			"BRL",
			basepb.Currency_CURRENCY_BRL,
		},
		{
			"Should return `CAD` currency",
			"CAD",
			basepb.Currency_CURRENCY_CAD,
		},
		{
			"Should return `USD` currency",
			"USD",
			basepb.Currency_CURRENCY_USD,
		},
		{
			"Should return `UNDEFINED` currency",
			"fake-currency",
			basepb.Currency_CURRENCY_UNDEFINED,
		},
	}
	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := NewFruitParser().ParseCurrency(tc.currency)
			assert.Equal(t, tc.response, resp)
			t.Logf("Parsed country %q : %v", tc.currency, resp)
		})
	}
}
