package kthlargestelement

import "testing"

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		kthLargest          int
		expectedValue       int
		expectedErrorString string
	}{
		{
			name:          "7 is 5th largest value in sorted array",
			intArray:      []int{0, 1, 2, 3, 5, 7, 8, 9, 11, 13},
			kthLargest:    5,
			expectedValue: 7,
		},
		{
			name:          "9 is 4th largest value in unsorted array",
			intArray:      []int{7, 92, 23, 9, -1, 0, 11, 6},
			kthLargest:    4,
			expectedValue: 9,
		},
		{
			name:                "0-th largest value throws correct error",
			intArray:            []int{7, 92, 23, 9, -1, 0, 11, 6},
			kthLargest:          0,
			expectedErrorString: "Cannot find 0-th largest value in array",
		},
		{
			name:                "Empty array throws correct error",
			intArray:            []int{},
			expectedErrorString: "Error finding kth largest of array: length of array is 0",
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

			received := kthLargest(test.intArray, test.kthLargest)

			if received != test.expectedValue {
				t.Errorf("Expected %d, received %d instead.", test.expectedValue, received)
			}
		})
	}
}

func TestRandomizedSelect(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		kthLargest          int
		expectedValue       int
		expectedErrorString string
	}{
		{
			name:          "7 is 5th largest value in sorted array",
			intArray:      []int{0, 1, 2, 3, 5, 7, 8, 9, 11, 13},
			kthLargest:    5,
			expectedValue: 7,
		},
		{
			name:          "9 is 4th largest value in unsorted array",
			intArray:      []int{7, 92, 23, 9, -1, 0, 11, 6},
			kthLargest:    4,
			expectedValue: 9,
		},
		{
			name:                "0-th largest value throws correct error",
			intArray:            []int{7, 92, 23, 9, -1, 0, 11, 6},
			kthLargest:          0,
			expectedErrorString: "Cannot find 0-th largest value in array",
		},
		{
			name:                "Empty array throws correct error",
			intArray:            []int{},
			expectedErrorString: "Error finding kth largest of array: length of array is 0",
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

			received := randomizedSelect(test.intArray, test.kthLargest)

			if received != test.expectedValue {
				t.Errorf("Expected %d, received %d instead.", test.expectedValue, received)
			}
		})
	}
}
