package kthlargestelement

import (
	"math/rand"
	"sort"
)

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

func randomizedSelect(array []int, k int) int {
	if len(array) == 0 {
		panic("Error finding kth largest of array: length of array is 0")
	}

	if k == 0 {
		panic("Cannot find 0-th largest value in array")
	}

	randomPivot := func(a []int, low int, high int) int {
		pivotIndex := rand.Intn(high-low) + low
		a[pivotIndex], a[high] = a[high], a[pivotIndex]
		return a[high]
	}

	randomizedPartition := func(a []int, low int, high int) int {
		pivot := randomPivot(a, low, high)
		i := low
		for j := low; j < high; j++ {
			if a[j] <= pivot {
				a[i], a[j] = a[j], a[i]
				i++
			}
		}
		a[i], a[high] = a[high], a[i]
		return i
	}

	var randomizedSelect func(a []int, low int, high int, k int) int
	randomizedSelect = func(a []int, low int, high int, k int) int {
		if low < high {
			p := randomizedPartition(a, low, high)
			if k == p {
				return a[p]
			} else if k < p {
				return randomizedSelect(a, low, p-1, k)
			} else {
				return randomizedSelect(a, p+1, high, k)
			}
		} else {
			return a[low]
		}
	}

	return randomizedSelect(array, 0, len(array)-1, k)
}
