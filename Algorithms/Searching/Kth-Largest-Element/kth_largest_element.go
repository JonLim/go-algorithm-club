package kthlargestelement

import "sort"

func kthLargest(array []int, k int) int {
	if len(array) == 0 {
		panic("Error finding kth largest of array: length of array is 0")
	}

	if k == 0 {
		panic("Cannot find 0-th largest value in array")
	}

	// Sorts `array` in ascending order
	sort.IntSlice(array).Sort()

	return array[len(array)-k]
}
