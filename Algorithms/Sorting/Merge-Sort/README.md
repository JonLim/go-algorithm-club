# Merge Sort

**Goal:** Sort an array from low to high (or high to low)

Invented in 1945 by John von Neumann, merge-sort is an efficient algorithm with a best, worst, and average time complexity of **O(n log n)**.

The merge-sort algorithm uses the **divide and conquer** approach which is to divide a big problem into smaller problems and solve them. I think of the merge-sort algorithm as **split first** and **merge after**.

Assume you need to sort an array of *n* numbers in the right order. The merge-sort algorithm works as follows:

- Put the numbers in an unsorted pile.
- Split the pile into two. Now, you have **two unsorted piles** of numbers.
- Keep splitting the resulting piles until you cannot split anymore. In the end, you will have *n* piles with one number in each pile.
- Begin to **merge** the piles together by pairing them sequentially. During each merge, put the contents in sorted order. This is fairly easy because each individual pile is already sorted.

## An example

### Splitting

Assume you are given an array of *n* numbers as`[2, 1, 5, 4, 9]`. This is an unsorted pile. The goal is to keep splitting the pile until you cannot split anymore.

First, split the array into two halves: `[2, 1]` and `[5, 4, 9]`. Can you keep splitting them? Yes, you can!

Focus on the left pile. Split`[2, 1]` into `[2]` and `[1]`. Can you keep splitting them? No. Time to check the other pile.

Split `[5, 4, 9]` into `[5]` and `[4, 9]`. Unsurprisingly, `[5]` cannot be split anymore, but `[4, 9]` can be split into `[4]` and `[9]`.

The splitting process ends with the following piles: `[2]` `[1]` `[5]` `[4]` `[9]`. Notice that each pile consists of just one element.

### Merging

Now that you have split the array, you should **merge** the piles together **while sorting them**. Remember, the idea is to solve many small problems rather than a big one. For each merge iteration, you must be concerned at merging one pile with another.

Given the piles `[2]` `[1]` `[5]` `[4]` `[9]`, the first pass will result in `[1, 2]` and `[4, 5]` and `[9]`. Since `[9]` is the odd one out, you cannot merge it with anything during this pass.

The next pass will merge `[1, 2]` and `[4, 5]` together. This results in `[1, 2, 4, 5]`, with the `[9]` left out again because it is the odd one out.

You are left with only two piles `[1, 2, 4, 5]` and `[9]`, finally gets its chance to merge, resulting in the sorted array as `[1, 2, 4, 5, 9]`.

## Top-down implementation

Here's what merge sort may look like in Go:

```golang
func mergeSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	middleIndex := len(array) / 2
	leftArray := mergeSort(array[0:middleIndex])
	rightArray := mergeSort(array[middleIndex:len(array)-1])

	return merge(leftArray, rightArray)
}
```

A step-by-step explanation of how the code works:

1. If the array is empty or contains a single element, there is no way to split it into smaller pieces. You must just return the array.

2. Find the middle index.

3. Using the middle index from the previous step, recursively split the left side of the array.

4. Also, recursively split the right side of the array.

5. Finally, merge all the values together, making sure that it is always sorted.

Here's the merging algorithm:

```golang
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
```

This method may look scary, but it is quite straightforward:

1. You need two indexes to keep track of your progress for the two arrays while merging.

2. This is the merged array. It is empty right now, but you will build it up in subsequent steps by appending elements from the other arrays. Since you already know number of elements that will end up in this array, you reserve capacity to avoid reallocation overhead later.

3. This while-loop will compare the elements from the left and right sides and append them into the `orderedArr` while making sure that the result stays in order.

4. If control exits from the previous while-loop, it means that either the `leftArray` or the `rightArray` has its contents completely merged into the `orderedArr`. At this point, you no longer need to do comparisons. Just append the rest of the contents of the other array until there is no more to append.

As an example of how `merge()` works, suppose that we have the following piles: `leftArray = [1, 7, 8]` and `rightArray = [3, 6, 9]`. Note that each of these piles is individually sorted already -- that is always true with merge sort. These are merged into one larger sorted pile in the following steps:

	leftArray      rightArray      orderedArr
	[ 1, 7, 8 ]    [ 3, 6, 9 ]     [ ]
      l              r

The left index, here represented as `l`, points at the first item from the left pile, `1`. The right index, `r`, points at `3`. Therefore, the first item we add to `orderedArr` is `1`. We also move the left index `l` to the next item.

	leftArray      rightArray      orderedArr
	[ 1, 7, 8 ]    [ 3, 6, 9 ]     [ 1 ]
      -->l           r

Now `l` points at `7` but `r` is still at `3`. We add the smallest item to the ordered pile, so that is `3`. The situation is now:

	leftArray      rightArray      orderedArr
	[ 1, 7, 8 ]    [ 3, 6, 9 ]     [ 1, 3 ]
         l           -->r

This process repeats. At each step, we pick the smallest item from either the `leftArray` or the `rightArray` and add the item into the `orderedArr`:

	leftArray      rightArray      orderedArr
	[ 1, 7, 8 ]    [ 3, 6, 9 ]     [ 1, 3, 6 ]
         l              -->r

	leftArray      rightArray      orderedArr
	[ 1, 7, 8 ]    [ 3, 6, 9 ]     [ 1, 3, 6, 7 ]
         -->l              r

	leftArray      rightArray      orderedArr
	[ 1, 7, 8 ]    [ 3, 6, 9 ]     [ 1, 3, 6, 7, 8 ]
            -->l           r

Now, there are no more items in the left pile. We simply add the remaining items from the right pile, and we are done. The merged pile is `[ 1, 3, 6, 7, 8, 9 ]`.

Notice that, this algorithm is very simple: it moves from left-to-right through the two piles and at every step picks the smallest item. This works because we guarantee that each of the piles is already sorted.

## Bottom-up implementation

The implementation of the merge-sort algorithm you have seen so far is called the "top-down" approach because it first splits the array into smaller piles and then merges them. When sorting an array (as opposed to, say, a linked list) you can actually skip the splitting step and immediately start merging the individual array elements. This is called the "bottom-up" approach.

Here is a complete bottom-up implementation in Go:

```golang
func mergeSortBottomUp(array []int) []int {
	n := len(array)

	z := [][]int{array, array}
	d := 0

	width := 1

	for width < n {
		i := 0
		for i < n {
			j, l, r := i, i, i+width

			lmax := mathint.Min(l+width, n)
			rmax := mathint.Min(r+width, n)

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
```

It looks a lot more intimidating than the top-down version, but notice that the main body includes the same three `while` loops from `merge()`.

Notable points:

1. The Merge-sort algorithm needs a temporary working array because you cannot merge the left and right piles and at the same time overwrite their contents. Because allocating a new array for each merge is wasteful, we are using two working arrays, and we will switch between them using the value of `d`, which is either 0 or 1. The array `z[d]` is used for reading, and `z[1 - d]` is used for writing. This is called *double-buffering*.

2. Conceptually, the bottom-up version works the same way as the top-down version. First, it merges small piles of one element each, then it merges piles of two elements each, then piles of four elements each, and so on. The size of the pile is given by `width`. Initially, `width` is `1` but at the end of each loop iteration, we multiply it by two, so this outer loop determines the size of the piles being merged, and the subarrays to merge become larger in each step.

3. The inner loop steps through the piles and merges each pair of piles into a larger one. The result is written in the array given by `z[1 - d]`.

4. This is the same logic as in the top-down version. The main difference is that we're using double-buffering, so values are read from `z[d]` and written into `z[1 - d]`.

5. At this point, the piles of size `width` from array `z[d]` have been merged into larger piles of size `width * 2` in array `z[1 - d]`. Here, we swap the active array, so that in the next step we'll read from the new piles we have just created.

Example of how to use it:

```golang
array := []int{2, 1, 5, 4, 9}
mergeSortBottomUp(array)   // [1, 2, 4, 5, 9]
```

## Performance

The speed of the merge-sort algorithm is dependent on the size of the array it needs to sort. The larger the array, the more work it needs to do.

Whether or not the initial array is sorted already does not affect the speed of the merge-sort algorithm since you will be doing the same amount splits and comparisons regardless of the initial order of the elements.

Therefore, the time complexity for the best, worst, and average case will always be **O(n log n)**.

A disadvantage of the merge-sort algorithm is that it needs a temporary "working" array equal in size to the array being sorted. It is not an **in-place** sort, unlike for example [quicksort](../Quicksort/).

Most implementations of the merge-sort algorithm produce a **stable** sort. This means that array elements that have identical sort keys will stay in the same order relative to each other after sorting. This is not important for simple values such as numbers or strings, but it can be an issue when sorting more complex objects.

## See also

[Merge sort on Wikipedia](https://en.wikipedia.org/wiki/Merge_sort)

*Reference material written for Swift Algorithm Club by Kelvin Lau. Additions by Matthijs Hollemans.*\
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*
