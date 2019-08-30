package quicksort

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Quicksort function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Quicksort function correctly handles single element array",
			intArray:      []int{10},
			expectedArray: []int{10},
		},
		{
			name:          "Empty array returns empty array",
			intArray:      []int{},
			expectedArray: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := quicksort(test.intArray)

			if !reflect.DeepEqual(test.expectedArray, received) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, received)
			}
		})
	}
}

func TestParitionLomuto(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		expectedPos         int
		expectedErrorString string
	}{
		{
			name:        "Partition Lomuto function correctly return partition index",
			intArray:    []int{10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1, 8},
			expectedPos: 7,
		},
		{
			name:        "Partition Lomuto function correctly return partition index for single element array",
			intArray:    []int{10},
			expectedPos: 0,
		},
		{
			name:                "Empty array returns empty array",
			intArray:            []int{},
			expectedErrorString: "Error finding partition index: length of array is 0",
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

			received := partitionLomuto(test.intArray, 0, len(test.intArray)-1)

			if test.expectedPos != received {
				t.Errorf("Expected parition index %v, received partition index %v instead.", test.expectedPos, received)
			}
		})
	}
}

func TestQuicksortLomuto(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Quicksort Lomuto function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Quicksort Lomuto function correctly handles single element array",
			intArray:      []int{10},
			expectedArray: []int{10},
		},
		{
			name:          "Empty array returns empty array",
			intArray:      []int{},
			expectedArray: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quicksortLomuto(test.intArray, 0, len(test.intArray)-1)

			if !reflect.DeepEqual(test.expectedArray, test.intArray) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, test.intArray)
			}
		})
	}
}

func TestParitionHoare(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		expectedPos         int
		expectedErrorString string
	}{
		{
			name:        "Partition Hoare function correctly return partition index",
			intArray:    []int{10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1, 8},
			expectedPos: 8,
		},
		{
			name:        "Partition Hoare function correctly return partition index for single element array",
			intArray:    []int{10},
			expectedPos: 0,
		},
		{
			name:                "Empty array returns empty array",
			intArray:            []int{},
			expectedErrorString: "Error finding partition index: length of array is 0",
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

			received := partitionHoare(test.intArray, 0, len(test.intArray)-1)

			if test.expectedPos != received {
				t.Errorf("Expected parition index %v, received partition index %v instead.", test.expectedPos, received)
			}
		})
	}
}

func TestQuicksortHoare(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Quicksort Hoare function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Quicksort Hoare function correctly handles single element array",
			intArray:      []int{10},
			expectedArray: []int{10},
		},
		{
			name:          "Empty array returns empty array",
			intArray:      []int{},
			expectedArray: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quicksortHoare(test.intArray, 0, len(test.intArray)-1)

			if !reflect.DeepEqual(test.expectedArray, test.intArray) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, test.intArray)
			}
		})
	}
}

func TestQuicksortRandom(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Quicksort Random function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Quicksort Random function correctly handles single element array",
			intArray:      []int{10},
			expectedArray: []int{10},
		},
		{
			name:          "Empty array returns empty array",
			intArray:      []int{},
			expectedArray: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quicksortRandom(test.intArray, 0, len(test.intArray)-1)

			if !reflect.DeepEqual(test.expectedArray, test.intArray) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, test.intArray)
			}
		})
	}
}

func TestParitionDutchFlag(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		pivotIndex          int
		expectedSmaller     int
		expectedLarger      int
		expectedErrorString string
	}{
		{
			name:            "Partition Dutch Flag function correctly return partition index",
			intArray:        []int{10, 0, 3, 9, 2, 14, 26, 27, 1, 5, 8, -1, 8},
			pivotIndex:      5,
			expectedSmaller: 10,
			expectedLarger:  10,
		},
		{
			name:            "Partition Dutch Flag function correctly return partition index for single element array",
			intArray:        []int{10},
			pivotIndex:      0,
			expectedSmaller: 0,
			expectedLarger:  0,
		},
		{
			name:                "Empty array returns empty array",
			intArray:            []int{},
			expectedErrorString: "Error finding partition indices: length of array is 0",
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

			tup1, tup2 := partitionDutchFlag(test.intArray, 0, len(test.intArray)-1, test.pivotIndex)

			if test.expectedSmaller != tup1 && test.expectedLarger != tup2 {
				t.Errorf("Expected parition indices (%v, %v), received partition indices (%v, %v) instead.", test.expectedSmaller, test.expectedLarger, tup1, tup2)
			}
		})
	}
}

func TestQuickSortDutchFlag(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Quicksort Dutch Flag function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Quicksort Dutch Flag function correctly handles single element array",
			intArray:      []int{10},
			expectedArray: []int{10},
		},
		{
			name:          "Empty array returns empty array",
			intArray:      []int{},
			expectedArray: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quicksortDutchFlag(test.intArray, 0, len(test.intArray)-1)

			if !reflect.DeepEqual(test.expectedArray, test.intArray) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, test.intArray)
			}
		})
	}
}
