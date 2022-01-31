package repository

import (
	"fmt"
	"testing"

	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
)

// UNITS TEST ****************************

func TestReadFruitByID(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		id       int
		response *entity.Fruit
		err      error
	}{
		{
			"Success",
			"../../data/fruits-test.csv",
			2,
			&entity.Fruit{},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewFruitRepo(tc.filePath)
			fruit, err := repo.ReadFruitByID(tc.id)
			// assert.Equal(t, tc.err, err)
			t.Logf("Test Fruit: %+v", fruit)
			t.Logf("Test Error: %v", err)

		})
	}
}

func TestWriteFruit(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		fruit    *entity.Fruit
		err      error
	}{
		{
			"Should return success, no errors",
			"../../data/fruits-test-write.csv",
			&entity.Fruit{
				ID:    666,
				Name:  "Test fruit",
				Color: "Magenta",
			},
			nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewFruitRepo(tc.filePath)
			err := repo.WriteFruit(tc.fruit)
			t.Logf("Test Error: %v", err)
		})
	}
}

func TestReadAllFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		response *entity.Fruits
		err      string
	}{
		{
			"Should return success, no errors",
			"../../data/fruits-test.csv",
			&entity.Fruits{},
			"<nil>",
		},
		{
			"Should return error `no such file or directory`",
			"",
			nil,
			"open : no such file or directory",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewFruitRepo(tc.filePath)
			fruits, err := repo.ReadAllFruits()
			assert.Equal(t, tc.err, fmt.Sprintf("%v", err))
			if err == nil {
				t.Log("Total fruits:", len(*fruits))
				for _, f := range *fruits {
					t.Logf("%+v", f)
				}
			}
		})
	}
}

func TestParseFruit(t *testing.T) {
	var testCases = []struct {
		name        string
		recordParam []string
		reponse     *entity.Fruits
		reponseErrs []error
	}{
		{
			"success, no errors",
			[]string{"1", "test", "test description"},
			nil,
			[]error{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := &fruitRepo{}
			fruit, errs := repo.parseFruit(tc.recordParam)
			t.Log("ERRORS:", errs)
			// assert.Equal(t, tc.reponseErrs, errs)
			// assert.NotEmpty(t, tc.reponse, fruit)
			t.Log("ParsedFruit:", fruit)
		})
	}

}
