package main

import (
	"testing"
)

func TestBowlingScores(t *testing.T) {

	var testCases = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}},
			output: 60,
		},
		{
			input:  [][]int{{4, 6}, {4, 6}, {4, 6}, {4, 6}, {4, 6}, {4, 6}, {4, 6}, {4, 6}, {4, 6}, {4, 6, 4}},
			output: 140,
		},
		{
			input:  [][]int{{10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {0, 0}},
			output: 240,
		},
		{
			input:  [][]int{{10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 10, 10}},
			output: 300,
		},
		{
			input:  [][]int{{10, 0}, {3, 3}, {3, 3}, {3, 3}, {8, 2}, {2, 8}, {7, 3}, {3, 7}, {5, 5}, {10, 0, 10}},
			output: 131,
		},
	}

	for _, v := range testCases {
		got := bowlingScores(v.input)
		if got != v.output {
			t.Fail()
		}
	}
}
