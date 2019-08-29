package linearsearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		valueToFind   int
		expectedIndex int
	}{
		{
			name:          "Value to find exists at start of array",
			intArray:      []int{5, 2, 4, 7},
			valueToFind:   5,
			expectedIndex: 0,
		},
		{
			name:          "Value to find exists near start of array",
			intArray:      []int{5, 2, 4, 7},
			valueToFind:   2,
			expectedIndex: 1,
		},
		{
			name:          "Value to find exists at end of array",
			intArray:      []int{5, 2, 4, 7},
			valueToFind:   7,
			expectedIndex: 3,
		},
		{
			name:          "Value to find does not exist in array",
			intArray:      []int{5, 2, 4, 7},
			valueToFind:   9,
			expectedIndex: -1,
		},
		{
			name:          "Value to find exists at middle of larger array",
			intArray:      []int{5, 2, 4, 7, 10, 3, 8, 6, 1, 15, 22, 11, 9, 30, 12, 19, 38},
			valueToFind:   15,
			expectedIndex: 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := linearSearch(test.intArray, test.valueToFind)

			if received != test.expectedIndex {
				t.Errorf("Expected index %d, received %d instead.", test.expectedIndex, received)
			}
		})
	}
}
