package queue

import (
	"reflect"
	"testing"
)

func TestIntQueueIsEmpty(t *testing.T) {
	tests := []struct {
		name         string
		intQueue     intQueue
		expectedBool bool
	}{
		{
			name:         "IntQueue IsEmpty correctly identifies empty stack",
			intQueue:     intQueue{values: []int{}},
			expectedBool: true,
		},
		{
			name:         "IntQueue IsEmpty correctly identifies stack with values",
			intQueue:     intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			expectedBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intQueue.isEmpty()

			if !reflect.DeepEqual(test.expectedBool, received) {
				t.Errorf("Expected boolean value of %v, received boolean %v instead.", test.expectedBool, received)
			}
		})
	}
}

func TestIntQueueCount(t *testing.T) {
	tests := []struct {
		name          string
		intQueue      intQueue
		expectedCount int
	}{
		{
			name:          "IntQueue Count correctly identifies empty stack",
			intQueue:      intQueue{values: []int{}},
			expectedCount: 0,
		},
		{
			name:          "IntQueue Count correctly identifies stack with 12 values",
			intQueue:      intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			expectedCount: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intQueue.count()

			if !reflect.DeepEqual(test.expectedCount, received) {
				t.Errorf("Expected queue count value of %v, received value of %v instead.", test.expectedCount, received)
			}
		})
	}
}

func TestIntQueueEnqueue(t *testing.T) {
	tests := []struct {
		name           string
		intQueue       intQueue
		valueToEnqueue interface{}
		expectedQueue  intQueue
	}{
		{
			name:           "IntQueue Enqueue adds value to empty stack",
			intQueue:       intQueue{values: []int{}},
			valueToEnqueue: 3,
			expectedQueue:  intQueue{values: []int{3}},
		},
		{
			name:           "IntQueue Enqueue adds value to stack with existing values",
			intQueue:       intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			valueToEnqueue: 12,
			expectedQueue:  intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26, 12}},
		},
		{
			name:           "IntQueue Enqueue does not add non-int value to queue",
			intQueue:       intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
			valueToEnqueue: "a",
			expectedQueue:  intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.intQueue.enqueue(test.valueToEnqueue)

			if !reflect.DeepEqual(test.expectedQueue, test.intQueue) {
				t.Errorf("Expected queue %v, received queue %v instead.", test.expectedQueue, test.intQueue)
			}
		})
	}
}

func TestIntQueueDeQueue(t *testing.T) {
	tests := []struct {
		name          string
		intQueue      intQueue
		expectedValue interface{}
		expectedQueue intQueue
	}{
		{
			name:          "IntQueue Dequeue returns value of single element in queue, and reduces queue to empty",
			intQueue:      intQueue{values: []int{3}},
			expectedValue: 3,
			expectedQueue: intQueue{values: []int{}},
		},
		{
			name:          "IntQueue Dequeue returns value of single element in queue, and removes first value in queue",
			intQueue:      intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26, 12}},
			expectedValue: 10,
			expectedQueue: intQueue{values: []int{-1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26, 12}},
		},
		{
			name:          "IntQueue Dequeue returns nil on empty queue",
			intQueue:      intQueue{values: []int{}},
			expectedValue: nil,
			expectedQueue: intQueue{values: []int{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intQueue.dequeue()

			if !reflect.DeepEqual(test.expectedValue, received) {
				t.Errorf("Expected return value of %v, received value of %v instead.", test.expectedValue, received)
			}

			if !reflect.DeepEqual(test.expectedQueue, test.intQueue) {
				t.Errorf("Expected queue %v, received queue %v instead.", test.expectedQueue, test.intQueue)
			}
		})
	}
}

func TestIntQueueFront(t *testing.T) {
	tests := []struct {
		name          string
		intQueue      intQueue
		expectedValue interface{}
	}{
		{
			name:          "IntQueue Front returns value of first element in stack with single value",
			intQueue:      intQueue{values: []int{3}},
			expectedValue: 3,
		},
		{
			name:          "IntQueue Front returns value of first element in stack with many values",
			intQueue:      intQueue{values: []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26, 12}},
			expectedValue: 10,
		},
		{
			name:          "IntQueue Front returns nil on empty stack",
			intQueue:      intQueue{values: []int{}},
			expectedValue: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			received := test.intQueue.front()

			if !reflect.DeepEqual(test.expectedValue, received) {
				t.Errorf("Expected return value of %v, received value of %v instead.", test.expectedValue, received)
			}
		})
	}
}
