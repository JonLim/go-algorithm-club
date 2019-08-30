package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name          string
		intArray      []int
		expectedArray []int
	}{
		{
			name:          "Selection sort function correctly sorts array",
			intArray:      []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			expectedArray: []int{-1, 0, 1, 2, 3, 3, 5, 8, 9, 10, 26, 27},
		},
		{
			name:          "Selection sort function correctly handles single element array",
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
			received := selectionSort(test.intArray)

			if !reflect.DeepEqual(test.expectedArray, received) {
				t.Errorf("Expected sorted array %v, received array %v instead.", test.expectedArray, received)
			}
		})
	}
}
