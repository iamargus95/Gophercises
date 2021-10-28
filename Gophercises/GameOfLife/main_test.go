package main

import (
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {

	var testCases = []struct {
		desc   string
		input  [size][size]Cell
		output [size][size]Cell
	}{
		{
			desc:   "Simple 2 X 2 matrix",
			input:  [size][size]Cell{{{true, 0}, {false, 0}, {false, 0}}, {{true, 0}, {false, 0}, {true, 0}}, {{true, 0}, {false, 0}, {true, 0}}},
			output: [size][size]Cell{{{false, 1}, {true, 3}, {false, 1}}, {{true, 2}, {false, 5}, {false, 1}}, {{false, 1}, {false, 4}, {false, 1}}},
		},
		{
			desc:   "Simple 3 X 3 matrix",
			input:  [size][size]Cell{{{false, 0}, {true, 0}, {false, 0}}, {{true, 0}, {false, 0}, {true, 0}}, {{false, 0}, {false, 0}, {false, 0}}},
			output: [size][size]Cell{{{false, 2}, {true, 2}, {false, 2}}, {{false, 1}, {true, 3}, {false, 1}}, {{false, 1}, {false, 2}, {false, 1}}},
		},
	}

	for _, v := range testCases {

		input := updateAliveNeighbours(v.input)

		got := gameOfLife(input)
		if !reflect.DeepEqual(got, v.output) {
			t.Fail()
		}
	}
}
