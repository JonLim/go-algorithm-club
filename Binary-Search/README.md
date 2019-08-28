# Binary Search

**Goal:** Quickly find an element in an array.

Let's say you have an array of numbers and you want to determine whether a specific number is in that array, and if so, at which index.

In most cases, iterating over a `range` in Go is good enough for that:

```golang
numbers := []int{11, 59, 3, 2, 53, 17, 31, 7, 19, 67, 47, 13, 37, 61, 29, 43, 5, 41, 23}

for index := range numbers {
    if numbers[index] == 43 {
        return index // Returns 15
    }
}
```

However, this approach performs a [linear search](../Linear-Search/). In code, that looks something like this:

```golang
func linearSearch(nums []int, target int) int {
	for index := range nums {
		if nums[index] == target {
			return index
		}
	}

	return -1
}
```

And you'd use it like this:

```golang
linearSearch(numbers, 43)  // returns 15
```

So what's the problem? `linearSearch()` loops through the entire array from the beginning, until it finds the element you're looking for. In the worst case, the value isn't even in the array and all that work is done for nothing.

On average, the linear search algorithm needs to look at half the values in the array. If your array is large enough, this starts to become very slow!

## Divide and conquer

The classic way to speed this up is to use a *binary search*. The trick is to keep splitting the array in half until the value is found.

For an array of size `n`, the performance is not **O(n)** as with linear search but only **O(log n)**. To put that in perspective, binary search on an array with 1,000,000 elements only takes about 20 steps to find what you're looking for, because `log_2(1,000,000) = 19.9`. And for an array with a billion elements it only takes 30 steps. (Then again, when was the last time you used an array with a billion items?)

Sounds great, but there is a downside to using binary search: *the array must be sorted.* In practice, this usually isn't a problem.

Here's how binary search works:

- Split the array in half and determine whether the thing you're looking for, known as the *search key*, is in the left half or in the right half.
- How do you determine in which half the search key is? This is why you sorted the array first, so you can do a simple `<` or `>` comparison.
- If the search key is in the left half, you repeat the process there: split the left half into two even smaller pieces and look in which piece the search key must lie. (Likewise for when it's the right half.)
- This repeats until the search key is found. If the array cannot be split up any further, you must regrettably conclude that the search key is not present in the array.

Now you know why it's called a "binary" search: in every step it splits the array into two halves. This process of *divide and conquer* is what allows it to quickly narrow down where the search key must be.

## The code

Here is a recursive implementation of binary search in Swift:

```golang
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
```

To try this out, copy the code to the Go Playground and do:

```golang
numbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}
binarySearch(numbers, 43) // Gives 13
```

Note that the `numbers` array is sorted. The binary search algorithm does not work otherwise!

The initial explanation was that binary search works by splitting the array in half, but we don't actually create two new arrays. Instead, we keep track of these splits using two integers, which represents the "boundaries" of where we are searching in the array. Initially, these boundaries cover the entire array, `0, len(nums)-1`, and as we split the array, the boundaries become smaller and smaller.

## Stepping through the example

It might be useful to look at how the algorithm works in detail.

The array from the above example consists of 19 numbers and looks like this when sorted:

	[ 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67 ]

We're trying to determine if the number `43` is in this array.

To split the array in half, we need to know the index of the object in the middle. That's determined by this line:

```golang
midIndex := lowerBound + ((upperBound - lowerBound) / 2)
```

Initially, the boundaries are `lowerBound = 0` and `upperBound = 18` (because `len(nums)-1` provides the index of the last number in our array!). Filling in these values, we find that `midIndex` is `0 + ((18 - 0) / 2) = 0 + 9 = 9`.

In the next figure, the `*` shows the middle item. As you can see, the number of items on each side is the same, so we're split right down the middle.

	[ 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67 ]
                                      *

Now binary search will determine which half to use. The relevant section from the code is:

```golang
if nums[midIndex] > target { // Is the target number in the left half?
    upperBound = midIndex
} else if nums[midIndex] < target { // Is the target number in the right half?
    lowerBound = midIndex + 1
} else { // If we get here, then we've found the index of the target number!
    return midIndex
}
```

In this case, `nums[midIndex] = 29`. That's less than the target number, so we can safely conclude that the index will never be in the left half of the array. After all, the left half only contains numbers smaller than `29`. Hence, the index must be in the right half somewhere (or not in the array at all).

Now we can simply repeat the binary search, but on the array interval from `midIndex + 1` to `upperBound`, or `len(nums)-1`:

	[ x, x, x, x, x, x, x, x, x, x | 31, 37, 41, 43, 47, 53, 59, 61, 67 ]

Since we no longer need to concern ourselves with the left half of the array, I've marked that with `x`'s. From now on we'll only look at the right half, which starts at array index 10.

We calculate the index of the new middle element: `midIndex = 10 + ((18 - 10)/2) = 10 + 4 = 14`, and split the array down the middle again.

	[ x, x, x, x, x, x, x, x, x, x | 31, 37, 41, 43, 47, 53, 59, 61, 67 ]
	                                                 *

As you can see, `nums[14]` is indeed the middle element of the array's right half.

Is the search key greater or smaller than `nums[14]`? It's smaller because `43 < 47`. This time we're taking the left half and ignore the larger numbers on the right:

	[ x, x, x, x, x, x, x, x, x, x | 31, 37, 41, 43 | x, x, x, x, x ]

The new `midIndex` is here (`10 + ((14-10)/2) = 10 + 2 = 12`):

	[ x, x, x, x, x, x, x, x, x, x | 31, 37, 41, 43 | x, x, x, x, x ]
                                             *

The search key is greater than `41`, so continue with the right side:

	[ x, x, x, x, x, x, x, x, x, x | x, x | x | 43 | x, x, x, x, x ]
	                                            *

And now we're done. The search key equals the array element we're looking at, so we've finally found what we were searching for: number `43` is at array index `13`. w00t!

It may have seemed like a lot of work, but in reality it only took four steps to find the search key in the array, which sounds about right because `log_2(19) = 4.23`. With a linear search, it would have taken 14 steps.

What would happen if we were to search for `42` instead of `43`? In that case, we can't split up the array any further. The `range.upperBound` becomes smaller than `range.lowerBound`. That tells the algorithm the search key is not in the array and it returns `nil`.

> **Note:** Many implementations of binary search calculate `midIndex = (lowerBound + upperBound) / 2`. This contains a subtle bug that only appears with very large arrays, because `lowerBound + upperBound` may overflow the maximum number an integer can hold. This situation is unlikely to happen on a 64-bit CPU, but it definitely can on 32-bit machines.

## Iterative vs recursive

Binary search is recursive in nature because you apply the same logic over and over again to smaller and smaller subarrays. However, that does not mean you must implement `binarySearch()` as a recursive function. It's often more efficient to convert a recursive algorithm into an iterative version, using a simple loop instead of lots of recursive function calls.

Here is a recursive implementation of binary search in Go:

```golang
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
```

As you can see, the code is very similar to the iterative version. The main difference is the function having to keep track of the boundaries in the function signature itself.

Use it like this:

```golang
numbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}
binarySearchRecursive(numbers, 43, 0, len(numbers)-1)  // gives 13
```

## The end

Is it a problem that the array must be sorted first? It depends. Keep in mind that sorting takes time -- the combination of binary search plus sorting may be slower than doing a simple linear search. Binary search shines in situations where you sort just once and then do many searches.

See also [Wikipedia](https://en.wikipedia.org/wiki/Binary_search_algorithm).

*Reference material written for Swift Algorithm Club by Matthijs Hollemans*\
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*
