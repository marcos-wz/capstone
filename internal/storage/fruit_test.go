package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFruits(t *testing.T) {
	var testCases = []struct {
		name     string
		filePath string
		err      string
	}{
		{
			"Should return success, no errors",
			"../../data/fruits-test.csv",
			"<nil>",
		},
		{
			"Should return error `no such file or directory`",
			"",
			"open : no such file or directory",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Initial data
			stg := NewFruitStorage(tc.filePath)
			// Test unit
			fruits, err := stg.ReadFruits()
			assert.Equal(t, tc.err, fmt.Sprintf("%v", err))
			if err == nil {
				assert.NotEmpty(t, fruits)
				t.Logf("Total fruits: %d", len(fruits))
				for _, fruit := range fruits {
					t.Logf("%+v", fruit)
				}
			}
		})
	}
}
