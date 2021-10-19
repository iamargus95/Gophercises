package main

import "testing"

func TestCipher(t *testing.T) {

	var testCases = []struct {
		input    rune
		delta    int
		expected rune
	}{
		{
			input:    'a',
			delta:    1,
			expected: 'b',
		},
	}

	for _, v := range testCases {
		got := cipher(v.input, v.delta, alphabet)
		if got != v.expected {
			t.Fail()
		}
	}
}
