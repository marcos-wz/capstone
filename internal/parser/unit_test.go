package parser

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_ParseUnit(t *testing.T) {
	var testCases = []struct {
		name     string
		unit     string
		response basepb.Unit
	}{
		{
			"Should return `LB` unit",
			"LB",
			basepb.Unit_UNIT_LB,
		},
		{
			"Should return `KG` unit",
			"KG",
			basepb.Unit_UNIT_KG,
		},
		{
			"Should return `UNDEFINED` unit",
			"fake-unit",
			basepb.Unit_UNIT_UNDEFINED,
		},
	}
	// RUN TESTS --------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := NewFruitParser().ParseUnit(tc.unit)
			assert.Equal(t, tc.response, resp)
			t.Logf("Parsed unit %q : %v", tc.unit, resp)
		})
	}
}
