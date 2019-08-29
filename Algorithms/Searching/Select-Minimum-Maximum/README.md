# Select Minimum / Maximum

**Goal:** Find the minimum/maximum object in an unsorted array.

## Maximum or minimum

We have an array of generic objects and we iterate over all the objects keeping track of the minimum/maximum element so far.

### An example

Let's say the we want to find the maximum value in the unsorted list `[ 8, 3, 9, 4, 6 ]`.

Pick the first number, `8`, and store it as the maximum element so far. 

Pick the next number from the list, `3`, and compare it to the current maximum. `3` is less than `8` so the maximum `8` does not change.

Pick the next number from the list, `9`, and compare it to the current maximum. `9` is greater than `8` so we store `9` as the maximum.

Repeat this process until the all elements in the list have been processed.

### The code

Here is a simple implementation in Go:

```golang
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
```

Put this code in a playground and test it like so:

```golang
array := []int{8, 3, 9, 4, 6}
minimum(array) // This will return 3
maximum(array) // This will return 9
```

## Performance

This algorithm runs at **O(n)**. Each object in the array is compared with the running minimum/maximum so the time it takes is proportional to the array length.

*Reference material written for Swift Algorithm Club by [Chris Pilcher](https://github.com/chris-pilcher)*\
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*
