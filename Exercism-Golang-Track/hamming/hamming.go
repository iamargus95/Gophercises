// Package hamming has functions that help compare two strands of DNA
package hamming

import "fmt"

// Func Distance compares two strands of DNA only if they are of equal length and gives the number of differences.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("provided DNA strands are of different lengths")
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
