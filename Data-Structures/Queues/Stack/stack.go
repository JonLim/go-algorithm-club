package stack

type Stack interface {
	isEmpty() bool
	count() int
	push(element interface{})
	pop() interface{}
	top() interface{}
}

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
