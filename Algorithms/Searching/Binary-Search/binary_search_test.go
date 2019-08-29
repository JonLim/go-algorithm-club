package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		valueToFind   int
		expectedIndex int
	}{
		{
			name:          "Value to find exists at start of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   2,
			expectedIndex: 0,
		},
		{
			name:          "Value to find exists near start of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   3,
			expectedIndex: 1,
		},
		{
			name:          "Value to find exists at end of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   67,
			expectedIndex: 18,
		},
		{
			name:          "Value to find exists mear end of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   59,
			expectedIndex: 16,
		},
		{
			name:          "Value to find does not exist in array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   100,
			expectedIndex: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := binarySearch(test.intArray, test.valueToFind)

			if received != test.expectedIndex {
				t.Errorf("Expected index %d, received %d instead.", test.expectedIndex, received)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		valueToFind   int
		expectedIndex int
	}{
		{
			name:          "Value to find exists at start of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   2,
			expectedIndex: 0,
		},
		{
			name:          "Value to find exists near start of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   3,
			expectedIndex: 1,
		},
		{
			name:          "Value to find exists at end of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   67,
			expectedIndex: 18,
		},
		{
			name:          "Value to find exists mear end of array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   59,
			expectedIndex: 16,
		},
		{
			name:          "Value to find does not exist in array",
			intArray:      []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67},
			valueToFind:   100,
			expectedIndex: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := binarySearchRecursive(test.intArray, test.valueToFind, 0, len(test.intArray)-1)

			if received != test.expectedIndex {
				t.Errorf("Expected index %d, received %d instead.", test.expectedIndex, received)
			}
		})
	}
}
