package countoccurrences

func countOccurrences(target int, array []int) int {
	leftBoundary := func() int {
		low, high := 0, len(array)
		for low < high {
			midIndex := low + ((high - low) / 2)
			if array[midIndex] < target {
				low = midIndex + 1
			} else {
				high = midIndex
			}
		}
		return low
	}

	rightBoundary := func() int {
		low, high := 0, len(array)
		for low < high {
			midIndex := low + ((high - low) / 2)
			if array[midIndex] > target {
				high = midIndex
			} else {
				low = midIndex + 1
			}
		}
		return low
	}

	return rightBoundary() - leftBoundary()
}
