package linearsearch

func linearSearch(nums []int, target int) int {
	for index := range nums {
		if nums[index] == target {
			return index
		}
	}

	return -1
}
