# Insertion Sort

**Goal:** Sort an array from low to high (or high to low).

You are given an array of numbers and need to put them in the right order. The insertion sort algorithm works as follows:

- Put the numbers on a pile. This pile is unsorted.
- Pick a number from the pile. It doesn't really matter which one you pick, but it's easiest to pick from the top of the pile. 
- Insert this number into a new array. 
- Pick the next number from the unsorted pile and also insert that into the new array. It either goes before or after the first number you picked, so that now these two numbers are sorted.
- Again, pick the next number from the pile and insert it into the array in the proper sorted position.
- Keep doing this until there are no more numbers on the pile. You end up with an empty pile and an array that is sorted.

That's why this is called an "insertion" sort, because you take a number from the pile and insert it in the array in its proper sorted position. 

## An example

Let's say the numbers to sort are `[ 8, 3, 5, 4, 6 ]`. This is our unsorted pile.

Pick the first number, `8`, and insert it into the new array. There is nothing in that array yet, so that's easy. The sorted array is now `[ 8 ]` and the pile is `[ 3, 5, 4, 6 ]`.

Pick the next number from the pile, `3`, and insert it into the sorted array. It should go before the `8`, so the sorted array is now `[ 3, 8 ]` and the pile is reduced to `[ 5, 4, 6 ]`.

Pick the next number from the pile, `5`, and insert it into the sorted array. It goes in between the `3` and `8`. The sorted array is `[ 3, 5, 8 ]` and the pile is `[ 4, 6 ]`.

Repeat this process until the pile is empty.

## In-place sort

The above explanation makes it seem like you need two arrays: one for the unsorted pile and one that contains the numbers in sorted order.

But you can perform the insertion sort *in-place*, without having to create a separate array. You just keep track of which part of the array is sorted already and which part is the unsorted pile.

Initially, the array is `[ 8, 3, 5, 4, 6 ]`. The `|` bar shows where the sorted portion ends and the pile begins:

	[| 8, 3, 5, 4, 6 ]

This shows that the sorted portion is empty and the pile starts at `8`.

After processing the first number, we have:

	[ 8 | 3, 5, 4, 6 ]

The sorted portion is `[ 8 ]` and the pile is `[ 3, 5, 4, 6 ]`. The `|` bar has shifted one position to the right.

This is how the content of the array changes during the sort:

	[| 8, 3, 5, 4, 6 ]
	[ 8 | 3, 5, 4, 6 ]
	[ 3, 8 | 5, 4, 6 ]
	[ 3, 5, 8 | 4, 6 ]
	[ 3, 4, 5, 8 | 6 ]
	[ 3, 4, 5, 6, 8 |]

In each step, the `|` bar moves up one position. As you can see, the beginning of the array up to the `|` is always sorted. The pile shrinks by one and the sorted portion grows by one, until the pile is empty and there are no more unsorted numbers left.

## How to insert

At each step you pick the top-most number from the unsorted pile and insert it into the sorted portion of the array. You must put that number in the proper place so that the beginning of the array remains sorted. How does that work?

Let's say we've already done the first few elements and the array looks like this:

	[ 3, 5, 8 | 4, 6 ]

The next number to sort is `4`. We need to insert that into the sorted portion `[ 3, 5, 8 ]` somewhere. 

Here's one way to do this: Look at the previous element, `8`. 

	[ 3, 5, 8, 4 | 6 ]
	        ^
	        
Is this greater than `4`? Yes it is, so the `4` should come before the `8`. We swap these two numbers to get:

	[ 3, 5, 4, 8 | 6 ]
	        <-->
	       swapped

We're not done yet. The new previous element, `5`, is also greater than `4`. We also swap these two numbers:

	[ 3, 4, 5, 8 | 6 ]
	     <-->
	    swapped

Again, look at the previous element. Is `3` greater than `4`? No, it is not. That means we're done with number `4`. The beginning of the array is sorted again.

This was a description of the inner loop of the insertion sort algorithm, which you'll see in the next section. It inserts the number from the top of the pile into the sorted portion by swapping numbers.

## The code

Here is an implementation of insertion sort in Go:

```golang
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
```

Put this code in a playground and test it like so:

```golang
array := []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}
insertionSort(array)
```

Here is how the code works.

1. There are two loops inside this function. The outer loop looks at each of the elements in the array in turn; this is what picks the top-most number from the pile. The variable `x` is the index of where the sorted portion ends and the pile begins (the position of the `|` bar). Remember, at any given moment the beginning of the array -- from index 0 up to `x` -- is always sorted. The rest, from index `x` until the last element, is the unsorted pile.

2. The inner loop looks at the element at position `x`. This is the number at the top of the pile, and it may be smaller than any of the previous elements. The inner loop steps backwards through the sorted array; every time it finds a previous number that is larger, it swaps them. When the inner loop completes, the beginning of the array is sorted again, and the sorted portion has grown by one element.

> **Note:** The outer loop starts at index 1, not 0. Moving the very first element from the pile to the sorted portion doesn't actually change anything, so we might as well skip it. 

## Performance

Insertion sort is really fast if the array is already sorted. That sounds obvious, but this is not true for all search algorithms. In practice, a lot of data will already be largely -- if not entirely -- sorted and insertion sort works quite well in that case.

The worst-case and average case performance of insertion sort is **O(n^2)**. That's because there are two nested loops in this function. Other sort algorithms, such as quicksort and merge sort, have **O(n log n)** performance, which is faster on large inputs.

Insertion sort is actually very fast for sorting small arrays. Some standard libraries have sort functions that switch from a quicksort to insertion sort when the partition size is 10 or less.

## See also

[Insertion sort on Wikipedia](https://en.wikipedia.org/wiki/Insertion_sort)

*Reference material written for Swift Algorithm Club by Matthijs Hollemans*\
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*
