// Package ocr provides a function for recognizing ascii art numbers.
package ocr

import "strings"

var table = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func recognizeDigit(s string) string {
	if d, ok := table[s]; ok {
		return d
	}
	return "?"
}

func recognizeLine(line []string) string {
	var o string
	for i := 0; i < len(line[0]); i += 3 {
		o += recognizeDigit(line[0][i:i+3] + line[1][i:i+3] + line[2][i:i+3])
	}
	return o
}

// Recognize returns a slice of strings containing the lines of numbers detected
// in the provided input.
func Recognize(s string) []string {
	var o []string
	lines := strings.Split(s, "\n")
	for i := 1; i < len(lines); i += 4 {
		o = append(o, recognizeLine(lines[i:i+3]))
	}
	return o
}
