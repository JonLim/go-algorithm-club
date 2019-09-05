package selectionsampling

import "testing"

func TestSelectionSample(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		numberToSample      int
		expectedErrorString string
	}{
		{
			name:           "Array with 4 samples from provided array",
			intArray:       []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			numberToSample: 4,
		},
		{
			name:           "Array with 8 samples from provided larger array",
			intArray:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			numberToSample: 8,
		},
		{
			name:                "Attempting to sample a number higher than the length of array throws correct error message",
			intArray:            []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}, // Length == 12
			numberToSample:      13,
			expectedErrorString: "Error getting samples from array: length of array is less than number of samples to gather",
		},
		{
			name:                "Empty array throws correct error mesage",
			intArray:            []int{},
			expectedErrorString: "Error getting samples from array: length of array is less than number of samples to gather",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if r != test.expectedErrorString {
						t.Errorf("Expected error %s, received %s instead.", test.expectedErrorString, r)
					}
				}
			}()

			received := selectionSample(test.intArray, test.numberToSample)

			if len(received) != test.numberToSample {
				t.Errorf("Expected %d samples, received %d samples instead.", test.numberToSample, len(received))
			}

			// TODO: Add testing for verifying that sample values are found in initial array
		})
	}
}

func TestReservoirSample(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		numberToSample      int
		expectedErrorString string
	}{
		{
			name:           "Array with 4 samples from provided array",
			intArray:       []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			numberToSample: 4,
		},
		{
			name:           "Array with 8 samples from provided larger array",
			intArray:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			numberToSample: 8,
		},
		{
			name:                "Attempting to sample a number higher than the length of array throws correct error message",
			intArray:            []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}, // Length == 12
			numberToSample:      13,
			expectedErrorString: "Error getting samples from array: length of array is less than number of samples to gather",
		},
		{
			name:                "Empty array throws correct error mesage",
			intArray:            []int{},
			expectedErrorString: "Error getting samples from array: length of array is less than number of samples to gather",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if r != test.expectedErrorString {
						t.Errorf("Expected error %s, received %s instead.", test.expectedErrorString, r)
					}
				}
			}()

			received := reservoirSample(test.intArray, test.numberToSample)

			if len(received) != test.numberToSample {
				t.Errorf("Expected %d samples, received %d samples instead.", test.numberToSample, len(received))
			}

			// TODO: Add testing for verifying that sample values are found in initial array
		})
	}
}

func TestSelectionSampleOrdered(t *testing.T) {
	tests := []struct {
		name                string
		intArray            []int
		numberToSample      int
		expectedErrorString string
	}{
		{
			name:           "Array with 4 samples from provided array",
			intArray:       []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26},
			numberToSample: 4,
		},
		{
			name:           "Array with 8 samples from provided larger array",
			intArray:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			numberToSample: 8,
		},
		{
			name:                "Attempting to sample a number higher than the length of array throws correct error message",
			intArray:            []int{10, -1, 3, 9, 2, 27, 8, 5, 1, 3, 0, 26}, // Length == 12
			numberToSample:      13,
			expectedErrorString: "Error getting samples from array: length of array is less than number of samples to gather",
		},
		{
			name:                "Empty array throws correct error mesage",
			intArray:            []int{},
			expectedErrorString: "Error getting samples from array: length of array is less than number of samples to gather",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if r != test.expectedErrorString {
						t.Errorf("Expected error %s, received %s instead.", test.expectedErrorString, r)
					}
				}
			}()

			received := selectionSampleOrdered(test.intArray, test.numberToSample)

			if len(received) != test.numberToSample {
				t.Errorf("Expected %d samples, received %d samples instead.", test.numberToSample, len(received))
			}

			// TODO: Add testing for verifying that sample values are found in initial array
		})
	}
}
