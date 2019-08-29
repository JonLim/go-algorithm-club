package countoccurrences

import "testing"

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		valueToFind         int
		expectedOccurrences int
	}{
		{
			name:                "Expect 4 occurrences of the number 3",
			intArray:            []int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11},
			valueToFind:         3,
			expectedOccurrences: 4,
		},
		{
			name:                "Expect 0 occurrences of the number 2",
			intArray:            []int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11},
			valueToFind:         2,
			expectedOccurrences: 0,
		},
		{
			name:                "Expect 1 occurrence of the number 8",
			intArray:            []int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11},
			valueToFind:         8,
			expectedOccurrences: 1,
		},
		{
			name:                "Expect 2 occurrence of the number 11",
			intArray:            []int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11},
			valueToFind:         11,
			expectedOccurrences: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := countOccurrences(test.valueToFind, test.intArray)

			if received != test.expectedOccurrences {
				t.Errorf("Expected index %d, received %d instead.", test.expectedOccurrences, received)
			}
		})
	}
}
