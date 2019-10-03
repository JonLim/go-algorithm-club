package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Merge sort function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Merge sort function correctly handles single element array",
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
			received := mergeSort(test.intArray)

			if !reflect.DeepEqual(test.expectedArray, received) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, received)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name          string
		leftArray     []int
		rightArray    []int
		expectedArray []int
	}{
		{
			name:          "Merge function correctly combines single element arrays",
			leftArray:     []int{10},
			rightArray:    []int{8},
			expectedArray: []int{8, 10},
		},
		{
			name:          "Merge function correctly combines single element arrays, with rightArray with greater val",
			leftArray:     []int{8},
			rightArray:    []int{10},
			expectedArray: []int{8, 10},
		},
		{
			name:          "Merge function correctly combines multi-element arrays",
			leftArray:     []int{8, 14},
			rightArray:    []int{10},
			expectedArray: []int{8, 10, 14},
		},
		{
			name:          "Merge function correctly combines larger multi-element arrays",
			leftArray:     []int{8, 14, 29},
			rightArray:    []int{10, 12, 42},
			expectedArray: []int{8, 10, 12, 14, 29, 42},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := merge(test.leftArray, test.rightArray)

			if !reflect.DeepEqual(test.expectedArray, received) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, received)
			}
		})
	}
}

func TestMergeSortBottomUp(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Merge sort bottom up function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Merge sort bottom up function correctly handles single element array",
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
			received := mergeSortBottomUp(test.intArray)

			if !reflect.DeepEqual(test.expectedArray, received) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, received)
			}
		})
	}
}
