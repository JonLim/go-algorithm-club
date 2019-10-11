package queue

type queue interface {
	isEmpty() bool
	count() int
	enqueue(element interface{})
	dequeue() interface{}
	front() interface{}
}

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
