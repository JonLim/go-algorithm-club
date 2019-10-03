package mergesort

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middleIndex := len(array) / 2
	leftArray := mergeSort(array[0:middleIndex])
	rightArray := mergeSort(array[middleIndex:len(array)])

	return merge(leftArray, rightArray)
}

func merge(leftArray, rightArray []int) []int {
	leftIndex, rightIndex := 0, 0
	orderedArr := []int{}

	for leftIndex < len(leftArray) && rightIndex < len(rightArray) {
		if leftArray[leftIndex] < rightArray[rightIndex] {
			orderedArr = append(orderedArr, leftArray[leftIndex])
			leftIndex++
		} else if leftArray[leftIndex] > rightArray[rightIndex] {
			orderedArr = append(orderedArr, rightArray[rightIndex])
			rightIndex++
		} else {
			orderedArr = append(orderedArr, leftArray[leftIndex])
			leftIndex++

			orderedArr = append(orderedArr, rightArray[rightIndex])
			rightIndex++
		}
	}

	for leftIndex < len(leftArray) {
		orderedArr = append(orderedArr, leftArray[leftIndex])
		leftIndex++
	}

	for rightIndex < len(rightArray) {
		orderedArr = append(orderedArr, rightArray[rightIndex])
		rightIndex++
	}

	return orderedArr
}

func mergeSortBottomUp(array []int) []int {
	n := len(array)

	z := [][]int{array, append([]int{}, array...)}
	d := 0

	width := 1

	for width < n {
		i := 0
		for i < n {
			j, l, r := i, i, i+width

			lmax := min(l+width, n)
			rmax := min(r+width, n)

			for l < lmax && r < rmax {
				if z[d][l] < z[d][r] {
					z[1-d][j] = z[d][l]
					l++
				} else {
					z[1-d][j] = z[d][r]
					r++
				}
				j++
			}

			for l < lmax {
				z[1-d][j] = z[d][l]
				j++
				l++
			}

			for r < rmax {
				z[1-d][j] = z[d][r]
				j++
				r++
			}

			i += width * 2
		}

		width = width * 2
		d = 1 - d
	}

	return z[d]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
