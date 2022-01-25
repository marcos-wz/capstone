package repository

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFruitByID(t *testing.T) {
	var testCases = []struct {
		name string
		id   int
		err  string
	}{
		{
			"Success",
			1,
			"<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewFruitRepo()
			fruit, err := repo.ReadFruitByID(tc.id)
			assert.EqualError(t, err, tc.err)
			log.Printf("TEST-Fruit: %+v", fruit)
			// log.Printf("TEST-Error: %v", err)

		})
	}
}
