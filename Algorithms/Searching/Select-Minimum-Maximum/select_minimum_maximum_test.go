package selectminimummaximum

import (
	"testing"
)

func TestMinimum(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		expectedMinimum     int
		expectedErrorString string
	}{
		{
			name:            "3 is minimum",
			intArray:        []int{8, 3, 9, 4, 6},
			expectedMinimum: 3,
		},
		{
			name:            "-6 is minimum",
			intArray:        []int{8, 3, 9, 4, 6, 7, -1, -6, -3},
			expectedMinimum: -6,
		},
		{
			name:                "Empty array",
			intArray:            []int{},
			expectedErrorString: "Error finding minimum of array: length of array is 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if r != test.expectedErrorString {
						t.Errorf("Expected error %s, received %s instead.", test.expectedErrorString, r)
					}
				}
			}()

			received := minimum(test.intArray)

			if received != test.expectedMinimum {
				t.Errorf("Expected index %d, received %d instead.", test.expectedMinimum, received)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		expectedMaximum     int
		expectedErrorString string
	}{
		{
			name:            "9 is maximum",
			intArray:        []int{8, 3, 9, 4, 6},
			expectedMaximum: 9,
		},
		{
			name:            "23 is minimum",
			intArray:        []int{8, 3, 9, 4, 6, 7, 23, -1, -6, -3},
			expectedMaximum: 23,
		},
		{
			name:                "Empty array",
			intArray:            []int{},
			expectedErrorString: "Error finding maximum of array: length of array is 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if r != test.expectedErrorString {
						t.Errorf("Expected error %s, received %s instead.", test.expectedErrorString, r)
					}
				}
			}()

			received := maximum(test.intArray)

			if received != test.expectedMaximum {
				t.Errorf("Expected index %d, received %d instead.", test.expectedMaximum, received)
			}
		})
	}
}
