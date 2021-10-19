package main

import "testing"

func TestCamelCase(t *testing.T) {

	var testCases = []struct {
		input    string
		expected int
	}{
		{
			input:    "atThe",
			expected: 2,
		},
		{
			input:    "hasBeenGood",
			expected: 3,
		},
		{
			input:    "one",
			expected: 1,
		},
	}

	for _, test := range testCases {
		got := camelCase(test.input)
		if got != test.expected {
			t.Fail()
		}
	}
}
