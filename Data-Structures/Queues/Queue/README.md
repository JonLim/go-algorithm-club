# Queue

A queue is a list where you can only insert new items at the back and remove items from the front. This ensures that the first item you enqueue is also the first item you dequeue. First come, first serve!

Why would you need this? Well, in many algorithms you want to add objects to a temporary list and pull them off this list later. Often the order in which you add and remove these objects matters.

A queue gives you a FIFO or first-in, first-out order. The element you inserted first is the first one to come out. It is only fair! (A similar data structure, the [stack](../Stack/), is LIFO or last-in first-out.)

For the sake of this explainer on Stacks, the stack we've created will be focused on being able to push and pop `int` values onto and out of the stack. Here's what it looks like:

```golang
type intQueue struct {
	values []int
}

func (i *intQueue) isEmpty() bool {
	return len(i.values) == 0
}

func (i *intQueue) count() int {
	return len(i.values)
}

func (i *intQueue) enqueue(element interface{}) {
	intElement, ok := element.(int)
	if ok {
		i.values = append(i.values, intElement)
	}
}

func (i *intQueue) dequeue() interface{} {
	if i.isEmpty() {
		return nil
	}

	firstElem := i.values[0]
	i.values = i.values[1:]
	return firstElem
}

func (i *intQueue) front() interface{} {
	if i.isEmpty() {
		return nil
	}

	return i.values[0]
}
```

Here is an example to enqueue a number:

```golang
queue.enqueue(10)
```

The queue is now `[ 10 ]`. Add the next number to the queue:

```golang
queue.enqueue(3)
```

The queue is now `[ 10, 3 ]`. Add one more number:

```golang
queue.enqueue(57)
```

The queue is now `[ 10, 3, 57 ]`. Let's dequeue to pull the first element off the front of the queue:

```golang
queue.dequeue()
```

This returns `10` because that was the first number we inserted. The queue is now `[ 3, 57 ]`. Everyone moved up by one place.

```golang
queue.dequeue()
```

This returns `3`, the next dequeue returns `57`, and so on. If the queue is empty, dequeuing returns `nil` or in some implementations it gives an error message.

> **Note:** A queue is not always the best choice. If the order in which the items are added and removed from the list is not important, you can use a [stack](../Stack/) instead of a queue. Stacks are simpler and faster.

> **Additional Note:** The memory allocated for the array in Go is never returned. For a long-living queue you should consider using dynamic data structures.

## See also

There are many ways to create a queue. Alternative implementations use a [linked list](../Linked-List/), a [circular buffer](../Ring-Buffer/), or a [heap](../Heap/). 

Variations on this theme are [deque](../Deque/), a double-ended queue where you can enqueue and dequeue at both ends, and [priority queue](../Priority-Queue/), a sorted queue where the "most important" item is always at the front.

*Reference material written for Swift Algorithm Club by Matthijs Hollemans*
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*