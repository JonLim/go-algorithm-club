package selectionsampling

import "math/rand"

func selectionSample(array []int, k int) []int {
	if len(array) < k {
		panic("Error getting samples from array: length of array is less than number of samples to gather")
	}

	for i := 0; i < k; i++ {
		r := rand.Intn((len(array)-1)-i) + i
		if i != r {
			array[i], array[r] = array[r], array[i]
		}
	}

	return array[0:k]
}

func reservoirSample(array []int, k int) []int {
	if len(array) < k {
		panic("Error getting samples from array: length of array is less than number of samples to gather")
	}

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, array[i])
	}

	for i := k; i < len(array)-1; i++ {
		j := rand.Intn(i)
		if j < k {
			result[j] = array[i]
		}
	}

	return result
}

func selectionSampleOrdered(array []int, requested int) []int {
	if len(array) < requested {
		panic("Error getting samples from array: length of array is less than number of samples to gather")
	}

	examined, selected := 0, 0
	b := []int{}

	for selected < requested {
		r := rand.Float64()

		leftToExamine := len(array) - examined
		leftToAdd := requested - selected

		if float64(leftToExamine)*r < float64(leftToAdd) {
			selected++
			b = append(b, array[examined])
		}

		examined++
	}

	return b
}
