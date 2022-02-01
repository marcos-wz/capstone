package repository

import (
	"fmt"
	"testing"

	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestReader_ReadFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		response []entity.Fruit
		err      string
	}{
		{
			"Should return valid fruit list",
			"../../data/fruits-test-ok.csv",
			[]entity.Fruit{},
			"<nil>",
		},
		{
			"Should return error `no such file or directory`",
			"",
			nil,
			"open : no such file or directory",
		},
		{
			"Should return valid fruit list, with parser error",
			"../../data/fruits-test-error.csv",
			[]entity.Fruit{},
			"parser error",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewReaderRepo(tc.filePath)
			fruits, err := repo.ReadFruits()
			assert.Equal(t, tc.err, fmt.Sprintf("%v", err))
			if err == nil {
				t.Log("Total fruits:", len(fruits))
				for _, f := range fruits {
					t.Logf("%+v", f)
				}
			}
		})
	}
}

func TestReader_ParseFruit(t *testing.T) {
	var testCases = []struct {
		name        string
		recordParam []string
		reponse     *entity.Fruit
		reponseErrs []error
	}{
		{
			"success, no errors",
			[]string{"1", "test", "test description"},
			&entity.Fruit{},
			[]error{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &readerRepo{}
			fruit, errs := r.parseFruitRecord(tc.recordParam)
			t.Log("ERRORS:", errs)
			// assert.Equal(t, tc.reponseErrs, errs)
			// assert.NotEmpty(t, tc.reponse, fruit)
			t.Log("ParsedFruit:", fruit)
		})
	}
}
