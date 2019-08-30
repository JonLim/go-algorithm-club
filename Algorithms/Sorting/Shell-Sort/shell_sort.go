package shellsort

func shellSort(array []int) []int {
	subsetGap := len(array) / 2
	for subsetGap > 0 {
		for pos := 0; pos <= subsetGap; pos++ {
			array = insertionSort(array, pos, subsetGap)
		}
		subsetGap = subsetGap / 2
	}

	return array
}

// insertionSort is a particular implementation used in this shell sort package,
// where it sorts an array in place, with a smaller subset as determined by the
// gap parameter.
func insertionSort(array []int, start int, gap int) []int {
	for ind := start + gap; ind < len(array); ind += gap {
		currVal := array[ind]
		pos := ind
		for pos >= gap && array[pos-gap] > currVal {
			array[pos] = array[pos-gap]
			pos -= gap
		}
		array[pos] = currVal
	}

	return array
}
