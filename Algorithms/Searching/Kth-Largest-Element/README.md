# k-th Largest Element Problem

**Goal:** Find the *k-th* largest element in an array.

We have an array of integers, and we want to find the k-th largest element in an array. For example, the 1-st largest element is the maximum value that occurs in the array. If the array has *n* elements, the *n*-th largest element is the minimum. The median is the *n/2*-th largest element.

## The naive solution

The following solution is semi-naive. Its time complexity is **O(n log n)** since it first sorts the array, and therefore also uses additional **O(n)** space.

```golang
import "sort"

func kthLargest(array []int, k int) int {
	if len(array) == 0 {
		panic("Error finding kth largest of array: length of array is 0")
	}

	if k == 0 {
		panic("Cannot find 0-th largest value in array")
	}

    // Sorts `array` in ascending order
    sort.IntSlice(array).Sort()
    
	return array[len(array)-k]
}
```

The `kthLargest()` function takes two parameters: the array `array` consisting of integers, and `k`. It returns the *k*-th largest element.

Let's take a look at an example and run through the algorithm to see how it works. Given `k = 4` and the array:

    [ 7, 92, 23, 9, -1, 0, 11, 6 ]

Initially there's no direct way to find the k-th largest element, but after sorting the array it's rather straightforward. Here's the sorted array:

    [ -1, 0, 6, 7, 9, 11, 23, 92 ]

Now, all we must do is take the value, from the sorted array, at index `len(array) - k`:

```golang
array[len(array)-k] = array[8-4] = array[4] = 9
```

Conversely, if you were looking for the k-th *smallest* element, you'd use `array[k-1]`.

<!-- CURRENTLY UNDER CONSTRUCTION -->

## A faster solution

There is a clever algorithm that combines the ideas of [binary search](../Binary-Search/) and [quicksort]() to arrive at an **O(n)** solution.

Let's remember that binary search splits a sorted array in half over and over again, to quickly narrow in on the value you're searching for. That's what we'll do here too.

Quicksort also splits up arrays. It uses partitioning to move all smaller values to the left of the pivot and all greater values to the right. After partitioning around a certain pivot, that pivot value will already be in its final, sorted position. We can use that to our advantage here.

Here's how it works: We choose a random pivot, partition the array around that pivot, then act like a binary search and only continue in the left or right partition. This repeats until we've found a pivot that happens to end up in the *k*-th position.

Let's look at the original example again. We're looking for the 4-th largest element in this array:

	[ 7, 92, 23, 9, -1, 0, 11, 6 ]

The algorithm is a bit easier to follow if we look for the k-th *smallest* item instead, so let's take `k = 4` and look for the 4-th smallest element.

Note that we don't have to sort the array first. We pick one of the elements at random to be the pivot, let's say `11`, and partition the array around that. We might end up with something like this:

	[ 7, 9, -1, 0, 6, 11, 92, 23 ]
	 <------ smaller    larger -->

As you can see, all values smaller than `11` are on the left; all values larger are on the right. The pivot value `11` is now in its final place. The index of the pivot is 5, so the 4-th smallest element must be in the left partition somewhere. We can ignore the rest of the array from now on:

	[ 7, 9, -1, 0, 6, x, x, x ]

Again, let's pick a random pivot, let's say `6`, and partition the array around it. We might end up with something like this:

	[ -1, 0, 6, 9, 7, x, x, x ]

Pivot `6` ended up at index 2, so the 4-th smallest item must be in the right partition. We can ignore the left partition:

	[ x, x, x, 9, 7, x, x, x ]

Again we pick a pivot value at random, let's say `9`, and partition the array:

	[ x, x, x, 7, 9, x, x, x ]

The index of pivot `9` is 4, and that's exactly the *k* we're looking for. We're done! Notice how this only took a few steps and we did not have to sort the array first.

The following function implements these ideas:

```swift
public func randomizedSelect<T: Comparable>(_ array: [T], order k: Int) -> T {
  var a = array

  func randomPivot<T: Comparable>(_ a: inout [T], _ low: Int, _ high: Int) -> T {
    let pivotIndex = random(min: low, max: high)
    a.swapAt(pivotIndex, high)
    return a[high]
  }

  func randomizedPartition<T: Comparable>(_ a: inout [T], _ low: Int, _ high: Int) -> Int {
    let pivot = randomPivot(&a, low, high)
    var i = low
    for j in low..<high {
      if a[j] <= pivot {
        a.swapAt(i, j)
        i += 1
      }
    }
    a.swapAt(i, high)
    return i
  }

  func randomizedSelect<T: Comparable>(_ a: inout [T], _ low: Int, _ high: Int, _ k: Int) -> T {
    if low < high {
      let p = randomizedPartition(&a, low, high)
      if k == p {
        return a[p]
      } else if k < p {
        return randomizedSelect(&a, low, p - 1, k)
      } else {
        return randomizedSelect(&a, p + 1, high, k)
      }
    } else {
      return a[low]
    }
  }

  precondition(a.count > 0)
  return randomizedSelect(&a, 0, a.count - 1, k)
}
```

To keep things readable, the functionality is split into three inner functions:

- `randomPivot()` picks a random number and puts it at the end of the current partition (this is a requirement of the Lomuto partitioning scheme, see the discussion on [quicksort](../Quicksort/) for more details).

- `randomizedPartition()` is Lomuto's partitioning scheme from quicksort. When this completes, the randomly chosen pivot is in its final sorted position in the array. It returns the array index of the pivot.

- `randomizedSelect()` does all the hard work. It first calls the partitioning function and then decides what to do next. If the index of the pivot is equal to the *k*-th number we're looking for, we're done. If `k` is less than the pivot index, it must be in the left partition and we'll recursively try again there. Likewise for when the *k*-th number must be in the right partition.

Pretty cool, huh? Normally quicksort is an **O(n log n)** algorithm, but because we only partition smaller and smaller slices of the array, the running time of `randomizedSelect()` works out to **O(n)**.

> **Note:** This function calculates the *k*-th smallest item in the array, where *k* starts at 0. If you want the *k*-th largest item, call it with `a.count - k`.

*Reference material written for Swift Algorithm Club by Daniel Speiser. Additions by Matthijs Hollemans.*\
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*