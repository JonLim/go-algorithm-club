package binarysearch

func binarySearch(nums []int, target int) int {
	// The boundaries of where we are searching in the array.
	lowerBound, upperBound := 0, len(nums)-1

	for lowerBound <= upperBound {
		// Calculate where to split the array.
		midIndex := lowerBound + ((upperBound - lowerBound) / 2)

		if nums[midIndex] > target { // Is the target number in the left half?
			upperBound = midIndex
		} else if nums[midIndex] < target { // Is the target number in the right half?
			lowerBound = midIndex + 1
		} else { // If we get here, then we've found the index of the target number!
			return midIndex
		}
	}

	// If we get here, then the target number is not present in the array.
	return -1
}

func binarySearchRecursive(nums []int, target int, lowerBound int, upperBound int) int {
	if lowerBound > upperBound {
		// If we get here, then the target number is not present in the array.
		return -1
	}

	// Calculate where to split the array.
	midIndex := lowerBound + ((upperBound - lowerBound) / 2)

	if nums[midIndex] > target { // Is the target number in the left half?
		return binarySearchRecursive(nums, target, lowerBound, midIndex)
	} else if nums[midIndex] < target { // Is the target number in the right half?
		return binarySearchRecursive(nums, target, midIndex+1, upperBound)
	} else { // If we get here, then we've found the index of the target number!
		return midIndex
	}
}
