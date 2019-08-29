package insertionsort

func insertionSort(array []int) []int {
	for ind := range array {
		indCopy := ind
		for indCopy > 0 && array[indCopy] < array[indCopy-1] {
			array[indCopy-1], array[indCopy] = array[indCopy], array[indCopy-1]
			indCopy--
		}
	}

	return array
}
