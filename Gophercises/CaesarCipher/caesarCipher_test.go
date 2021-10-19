package main

import "testing"

func TestCipher(t *testing.T) {

	var testCases = []struct {
		desc     string
		input    string
		delta    int
		expected string
	}{
		{
			desc:     "All Caps",
			input:    "ABC",
			delta:    3,
			expected: "DEF",
		},
		{
			desc:     "Mix of uppercase and lowercase",
			input:    "AbCd",
			delta:    1,
			expected: "BcDe",
		},
		{
			desc:     "Uppercase Edge case",
			input:    "XYZ",
			delta:    3,
			expected: "ABC",
		},
		{
			desc:     "Uppercase and lowercase mixed edge case",
			input:    "XyZaA",
			delta:    5,
			expected: "CdEfF",
		},
		{
			desc:     "Uppercase and Lowercase mix with not alphabetical value",
			input:    "Abc-Def",
			delta:    6,
			expected: "Ghi-Jkl",
		},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			got := getCipher(v.input, v.delta)
			if got != v.expected {
				t.Errorf("getCipher(%s,%d) = %s; want %s", v.input, v.delta, got, v.expected)
			}
		})
	}
}
