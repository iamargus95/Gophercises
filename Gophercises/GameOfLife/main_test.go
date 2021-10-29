package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {

	var testCases = []struct {
		desc   string
		input  [][]int
		output [][]int
	}{
		{
			desc:   "Simple 3 X 3 matrix",
			input:  [][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}},
			output: [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
		},
		{
			desc:   "Simple 4 X 4 matrix",
			input:  [][]int{{0, 1, 0, 1}, {1, 0, 1, 0}, {0, 1, 0, 1}, {1, 0, 1, 0}},
			output: [][]int{{0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}, {0, 1, 1, 0}},
		},
		{
			desc:   "Simple 5 X 5 matrix",
			input:  [][]int{{0, 0, 0, 0, 0}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 0, 0, 0, 0}},
			output: [][]int{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {1, 0, 0, 0, 1}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		},
	}

	for _, v := range testCases {

		neighbours := getAliveNeighbours(v.input)
		fmt.Println(neighbours)

		got := gameOfLife(v.input, neighbours)
		if !reflect.DeepEqual(got, v.output) {
			t.Fail()
		}
	}
}
