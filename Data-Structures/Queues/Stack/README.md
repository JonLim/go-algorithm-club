# Stack

A stack is like an array but with limited functionality. You can only *push* to add a new element to the top of the stack, *pop* to remove the element from the top, and *peek* at the top element without popping it off.

Why would you want to do this? Well, in many algorithms you want to add objects to a temporary list at some point and then pull them off this list again at a later time. Often the order in which you add and remove these objects matters.

A stack gives you a last-in first-out order, which is often abbreviated as **LIFO**. The element you pushed last is the first one to come off with the next pop. (A very similar data structure, the [queue](../Queue/), is first-in first-out, or FIFO.)

For the sake of this explainer on Stacks, the stack we've created will be focused on being able to push and pop `int` values onto and out of the stack. Here's what it looks like:

```golang
type intStack struct {
	values []int
}

func (i *intStack) isEmpty() bool {
	return len(i.values) == 0
}

func (i *intStack) count() int {
	return len(i.values)
}

func (i *intStack) push(element interface{}) {
	intElement, ok := element.(int)
	if ok {
		i.values = append(i.values, intElement)
	}
}

func (i *intStack) pop() interface{} {
	if len(i.values) == 0 {
		return nil
	}

	lastElem := i.values[len(i.values)-1]
	i.values = i.values[:len(i.values)-1]
	return lastElem
}

func (i *intStack) top() interface{} {
	if len(i.values) == 0 {
		return nil
	}
	return i.values[len(i.values)-1]
}
```

> NOTE: Check out the [unit tests for stacks](./stack_test.go) to see the above code, but also how we've created a `Stack` interface that we can use to build a stack for any type of data.

Let's create a Stack:

```golang
stack := intStack{} // This should create an `intStack` object with an empty `values` field
```

Let's push a number on the stack:

```golang
stack.push(10)
```

The stack is now `[ 10 ]`. Push the next number:

```golang
stack.push(3)
```

The stack is now `[ 10, 3 ]`. Push one more number:

```golang
stack.push(57)
```

The stack is now `[ 10, 3, 57 ]`. Let's pop the top number off the stack:

```golang
stack.pop()
```

This returns `57`, because that was the most recent number we pushed. The stack is `[ 10, 3 ]` again.

```golang
stack.top()
```

This returns `3`, which is the last item in the stack at the moment. If the stack is empty, popping returns `nil` or in some implementations it gives an error message ("stack underflow").

Notice that a push puts the new element at the end of the array, not the beginning. Inserting at the beginning of an array is expensive, an **O(n)** operation, because it requires all existing array elements to be shifted in memory. Adding at the end is **O(1)**; it always takes the same amount of time, regardless of the size of the array.

Fun fact about stacks: Each time you call a function or a method, the CPU places the return address on a stack. When the function ends, the CPU uses that return address to jump back to the caller. That's why if you call too many functions -- for example in a recursive function that never ends -- you get a so-called "stack overflow" as the CPU stack has run out of space.

*Reference material written for Swift Algorithm Club by Matthijs Hollemans*
*Adapted for Go Algorithm Club by [Jon Lim](https://github.com/JonLim)*