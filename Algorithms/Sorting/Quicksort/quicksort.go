package quicksort

import "math/rand"

func quicksort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	pivot := array[len(array)/2]

	less, equal, greater := []int{}, []int{}, []int{}
	for ind := range array {
		if array[ind] < pivot {
			less = append(less, array[ind])
		}
		if array[ind] == pivot {
			equal = append(equal, array[ind])
		}
		if array[ind] > pivot {
			greater = append(greater, array[ind])
		}
	}

	// Joining three arrays together; in other languages, this is the same as
	// `quicksort(less) + equal + quicksort(greater)`
	return append(append(quicksort(less), equal...), quicksort(greater)...)
}

func partitionLomuto(array []int, low int, high int) int {
	if len(array) == 0 {
		panic("Error finding partition index: length of array is 0")
	}

	pivot := array[high]

	i := low
	for j := low; j < high; j++ {
		if array[j] <= pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}

	array[i], array[high] = array[high], array[i]
	return i
}

func quicksortLomuto(array []int, low int, high int) {
	if low < high {
		p := partitionLomuto(array, low, high)
		quicksortLomuto(array, low, p-1)
		quicksortLomuto(array, p+1, high)
	}
}

func partitionHoare(array []int, low int, high int) int {
	if len(array) == 0 {
		panic("Error finding partition index: length of array is 0")
	}

	pivot := array[low]
	i, j := low-1, high+1

	for {
		for {
			j--
			if array[j] <= pivot {
				break
			}
		}
		for {
			i++
			if array[i] >= pivot {
				break
			}
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		} else {
			return j
		}
	}
}

func quicksortHoare(array []int, low int, high int) {
	if low < high {
		p := partitionHoare(array, low, high)
		quicksortHoare(array, low, p)
		quicksortHoare(array, p+1, high)
	}
}

func quicksortRandom(array []int, low int, high int) {
	if low < high {
		pivotIndex := rand.Intn(high-low) + low
		array[pivotIndex], array[high] = array[high], array[pivotIndex]

		p := partitionLomuto(array, low, high)
		quicksortRandom(array, low, p-1)
		quicksortRandom(array, p+1, high)
	}
}

func partitionDutchFlag(array []int, low int, high int, pivotIndex int) (int, int) {
	if len(array) == 0 {
		panic("Error finding partition indices: length of array is 0")
	}

	pivot := array[pivotIndex]
	smaller, equal, larger := low, low, high

	for equal <= larger {
		if array[equal] < pivot {
			array[smaller], array[equal] = array[equal], array[smaller]
			smaller++
			equal++
		} else if array[equal] == pivot {
			equal++
		} else {
			array[equal], array[larger] = array[larger], array[equal]
			larger--
		}
	}

	return smaller, larger
}

func quicksortDutchFlag(array []int, low int, high int) {
	if low < high {
		pivotIndex := rand.Intn(high-low) + low
		p, q := partitionDutchFlag(array, low, high, pivotIndex)
		quicksortDutchFlag(array, low, p-1)
		quicksortDutchFlag(array, q+1, high)
	}
}
