package main

import (
	"reflect"
	"testing"
)

func TestMatrixThis(t *testing.T) {

	var testCases = []struct {
		input    int
		expected [][]int
	}{
		{
			input:    2,
			expected: [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}},
		},
		{
			input:    3,
			expected: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}},
		},
		{
			input:    5,
			expected: [][]int{{5, 5, 5, 5, 5, 5, 5, 5, 5}, {5, 4, 4, 4, 4, 4, 4, 4, 5}, {5, 4, 3, 3, 3, 3, 3, 4, 5}, {5, 4, 3, 2, 2, 2, 3, 4, 5}, {5, 4, 3, 2, 1, 2, 3, 4, 5}, {5, 4, 3, 2, 2, 2, 3, 4, 5}, {5, 4, 3, 3, 3, 3, 3, 4, 5}, {5, 4, 4, 4, 4, 4, 4, 4, 5}, {5, 5, 5, 5, 5, 5, 5, 5, 5}},
		},
	}

	for _, v := range testCases {
		got := matrixThis(v.input)
		if !reflect.DeepEqual(got, v.expected) {
			t.Fail()
		}
	}
}
