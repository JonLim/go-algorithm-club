package stack

import (
	"reflect"
	"testing"
)

func TestIntStackIsEmpty(t *testing.T) {
	tests := []struct {
		name         string
		intStack     intStack
		expectedBool bool
	}{
		{
			name:         "IntStack IsEmpty correctly identifies empty stack",
			intStack:     intStack{values: []int{}},
			expectedBool: true,
		},
		{
			name:         "IntStack IsEmpty correctly identifies stack with values",
			intStack:     intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			expectedBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intStack.isEmpty()

			if !reflect.DeepEqual(test.expectedBool, received) {
				t.Errorf("Expected boolean value of %v, received boolean %v instead.", test.expectedBool, received)
			}
		})
	}
}

func TestIntStackCount(t *testing.T) {
	tests := []struct {
		name          string
		intStack      intStack
		expectedCount int
	}{
		{
			name:          "IntStack Count correctly identifies empty stack",
			intStack:      intStack{values: []int{}},
			expectedCount: 0,
		},
		{
			name:          "IntStack Count correctly identifies stack with 12 values",
			intStack:      intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			expectedCount: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intStack.count()

			if !reflect.DeepEqual(test.expectedCount, received) {
				t.Errorf("Expected stack count value of %v, received value of %v instead.", test.expectedCount, received)
			}
		})
	}
}

func TestIntStackPush(t *testing.T) {
	tests := []struct {
		name          string
		intStack      intStack
		valueToPush   interface{}
		expectedStack intStack
	}{
		{
			name:          "IntStack Push adds value to empty stack",
			intStack:      intStack{values: []int{}},
			valueToPush:   3,
			expectedStack: intStack{values: []int{3}},
		},
		{
			name:          "IntStack Push adds value to stack with existing values",
			intStack:      intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			valueToPush:   12,
			expectedStack: intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26, 12}},
		},
		{
			name:          "IntStack Push does not add non-int value to stack",
			intStack:      intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			valueToPush:   "a",
			expectedStack: intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.intStack.push(test.valueToPush)

			if !reflect.DeepEqual(test.expectedStack, test.intStack) {
				t.Errorf("Expected stack %v, received stack %v instead.", test.expectedStack, test.intStack)
			}
		})
	}
}

func TestIntStackPop(t *testing.T) {
	tests := []struct {
		name          string
		intStack      intStack
		expectedValue interface{}
		expectedStack intStack
	}{
		{
			name:          "IntStack Pop returns value of single element in stack, and reduces stack to empty",
			intStack:      intStack{values: []int{3}},
			expectedValue: 3,
			expectedStack: intStack{values: []int{}},
		},
		{
			name:          "IntStack Pop returns value of single element in stack, and removes last value in stack",
			intStack:      intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26, 12}},
			expectedValue: 12,
			expectedStack: intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
		},
		{
			name:          "IntStack Pop returns nil on empty stack",
			intStack:      intStack{values: []int{}},
			expectedValue: nil,
			expectedStack: intStack{values: []int{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intStack.pop()

			if !reflect.DeepEqual(test.expectedValue, received) {
				t.Errorf("Expected return value of %v, received value of %v instead.", test.expectedValue, received)
			}

			if !reflect.DeepEqual(test.expectedStack, test.intStack) {
				t.Errorf("Expected stack %v, received stack %v instead.", test.expectedStack, test.intStack)
			}
		})
	}
}

func TestIntStackTop(t *testing.T) {
	tests := []struct {
		name          string
		intStack      intStack
		expectedValue interface{}
	}{
		{
			name:          "IntStack Top returns value of last element in stack with single value",
			intStack:      intStack{values: []int{3}},
			expectedValue: 3,
		},
		{
			name:          "IntStack Top returns value of last element in stack with many values",
			intStack:      intStack{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26, 12}},
			expectedValue: 12,
		},
		{
			name:          "IntStack Top returns nil on empty stack",
			intStack:      intStack{values: []int{}},
			expectedValue: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intStack.top()

			if !reflect.DeepEqual(test.expectedValue, received) {
				t.Errorf("Expected return value of %v, received value of %v instead.", test.expectedValue, received)
			}
		})
	}
}
