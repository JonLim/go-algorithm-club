package selectionsort

func selectionSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	for ind := range array {
		lowestInd := ind

		for innerInd := ind + 1; innerInd < len(array); innerInd++ {
			if array[innerInd] < array[lowestInd] {
				lowestInd = innerInd
			}
		}

		if ind != lowestInd {
			array[ind], array[lowestInd] = array[lowestInd], array[ind]
		}
	}

	return array
}
