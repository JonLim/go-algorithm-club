package selectminimummaximum

import "math"

func minimum(array []int) int {
	if len(array) == 0 {
		panic("Error finding minimum of array: length of array is 0")
	}

	minimum := math.MaxInt8
	for index := range array {
		if array[index] < minimum {
			minimum = array[index]
		}
	}

	return minimum
}

func maximum(array []int) int {
	if len(array) == 0 {
		panic("Error finding maximum of array: length of array is 0")
	}

	maximum := math.MinInt8
	for index := range array {
		if array[index] > maximum {
			maximum = array[index]
		}
	}

	return maximum
}
