package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPP_Error(t *testing.T) {
	var testCases = []struct {
		name      string
		eType     ErrorType
		eTypeDesc string
		message   string
		err       *FruitError
	}{
		{
			"Should return undefined error",
			ErrUndefined,
			"undefined error",
			"test error message",
			&FruitError{ErrUndefined, "undefined error: test error message"},
		},
		{
			"Should return error csv file",
			ErrRepoFileCSV,
			"csv file error",
			"test error message",
			&FruitError{ErrRepoFileCSV, "csv file error: test error message"},
		},
		{
			"Should return error json file",
			ErrRepoFileJSON,
			"json file error",
			"test error message",
			&FruitError{ErrRepoFileJSON, "json file error: test error message"},
		},
		{
			"Should return error parse fruit",
			ErrParseFruitCSV,
			"parse fruit error",
			"test error message",
			&FruitError{ErrParseFruitCSV, "parse fruit error: test error message"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := NewFruitError(tc.eType, tc.message)
			switch err.(type) {
			case *FruitError:
				errFruit := err.(*FruitError)
				assert.Equal(t, tc.err.eType, errFruit.Type())
				assert.Equal(t, tc.eTypeDesc, errFruit.TypeDesc())
				assert.Equal(t, tc.err.message, errFruit.Error())
			default:
				t.Errorf("should be a fruit error type: %T", err)
			}
		})
	}
}
