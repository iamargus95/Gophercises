package main

import (
	"reflect"
	"testing"
)

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

func TestMatrixThis(t *testing.T) {

	for _, v := range testCases {
		got := matrixThis(v.input)
		if !reflect.DeepEqual(got, v.expected) {
			t.Fail()
		}
	}
}

func BenchmarkMatrix(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			matrixThis(v.input)
		}

	}
}
