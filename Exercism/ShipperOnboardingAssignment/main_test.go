package main

import "testing"

func TestLongestSubstring(t *testing.T) {

	var testCases = []struct {
		input  string
		output string
	}{
		{
			input:  "abcabcbb",
			output: "abc",
		},
		{
			input:  "bbbb",
			output: "b",
		},
		{
			input:  "pwwkew",
			output: "wke",
		},
		{
			input:  "",
			output: "",
		},
		{
			input:  "qweasfgaushasknafjsqwe",
			output: "knafjsqwe",
		},
	}

	for _, test := range testCases {
		got := longestSubstring(test.input)
		if got != test.output {
			t.Fail()
		}
	}
}
