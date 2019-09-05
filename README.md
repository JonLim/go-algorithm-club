# Welcome to the Go Algorithm Club!

Inspired by the [Swift Algorithm Club](https://github.com/raywenderlich/swift-algorithm-club), the aim is to showcase implementations of popular algorithms and data structures in [Go](https://golang.org/), with detailed explanations of how they work.

Personally speaking, the goal of this project is to **improve my understanding of how algorithms work**, primarily using Go for implementation, as that is what I work with, and choose to work with, in my daily programming work.

Overall, similarly to the Swift Algorithm Club, the goal is to **explain how algorithms work**, with a focus on clarity and readability of the code, not for a reusable library for your own projects. Much of the explanations will use the Swift Algorithm Club's readmes as a starting point, when possible.

Code is written with **Go 1.13**.

## Algorithms (WIP!)

### Searching

- [Linear Search](Algorithms/Searching/Linear-Search/). Find an element in an array.
- [Binary Search](Algorithms/Searching/Binary-Search/). Quickly find elements in a sorted array.
- [Count Occurrences](Algorithms/Searching/Count-Occurrences/). Count how often a value appears in an array.
- [Select Minimum / Maximum](Algorithms/Searching/Select-Minimum-Maximum/). Find the minimum/maximum value in an array.
- [k-th Largest Element](Algorithms/Searching/Kth-Largest-Element/). Find the k-th largest element in an array, such as the median.
- [Selection Sampling](Algorithms/Searching/Selection-Sampling/). Randomly choose a bunch of items from a collection.

### Sorting

#### Basic Sorts

- [Insertion Sort](Algorithms/Sorting/Insertion-Sort/). Sort values by inserting to correct position.
- [Selection Sort](Algorithms/Sorting/Selection-Sort/). Sort by comparing values.
- [Shell Sort](Algorithms/Sorting/Shell-Sort). Improvement of insertion sort, using smaller subsets.

#### Fast Sorts

- [Quicksort](Algorithms/Sorting/Quicksort/).
- [Merge Sort]().
- [Heap Sort]().

## Running Tests

Once you've cloned the directory locally, from the root directory og `go-algorithm-club`, you can run `go test ./...` in terminal and run all of the current tests!

Each algorithm and/or data structure has its own section in the directory, which can contain:

- A `README` file, which explains the theory of the algorithm, as well as using examples to illustrate how they work
- A Go file, which contains all of the code found in the `README`
- A Go test file, which has testing scenarios for all of the code found in the Go file

## Credits

Inspiration for the Go Algorithm Club comes from the [Swift Algorithm Club](https://github.com/raywenderlich/swift-algorithm-club), from the Ray Wenderlich community.

Go Algorithm Club was created by [Jon Lim](https://jonlim.ca/).

## License

All content is licensed under the terms of the MIT open source license.

By posting here, or by submitting any pull request through this forum, you agree that all content you submit or create, both code and text, is subject to this license. Jon Lim, and others will have all the rights described in the license regarding this content.  The precise terms of this license may be found [here](https://github.com/JonLim/go-algorithm-club/blob/master/LICENSE).
