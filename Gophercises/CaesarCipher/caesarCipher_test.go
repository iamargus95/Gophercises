package main

import "testing"

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
	{
		desc:     "UPPERCASE with delta 26",
		input:    "AZ",
		delta:    26,
		expected: "AZ",
	},
	{
		desc:     "LowerCase with Delta 26",
		input:    "az",
		delta:    26,
		expected: "az",
	},
}

func TestCipher(t *testing.T) {

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			got := getCipher(v.input, v.delta)
			if got != v.expected {
				t.Errorf("getCipher(%s,%d) = %s; want %s", v.input, v.delta, got, v.expected)
			}
		})
	}
}

func BenchmarkCipher(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			getCipher(v.input, v.delta)
		}

	}
}
