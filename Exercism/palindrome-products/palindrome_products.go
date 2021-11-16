// Package palindrome calculates palindrome factors.
package palindrome

import (
	"errors"
	"fmt"
)

// Product represents a palindrome and its factors.
type Product struct {
	Value          int
	Factorizations [][2]int
}

func reverse(s string) string {
	var result string
	for _, r := range s {
		result = string(r) + result
	}
	return result
}

// Products returns palindromes within the range and their factors.
func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	var smallest *Product
	var largest *Product
	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			p := x * y
			ps := fmt.Sprintf("%v", p)

			if ps != reverse(ps) {
				continue
			}

			if smallest == nil || p < smallest.Value {
				smallest = &Product{Value: p, Factorizations: [][2]int{{x, y}}}
			} else if smallest != nil && p == smallest.Value {
				smallest.Factorizations = append(smallest.Factorizations, [2]int{x, y})
			}

			if largest == nil || p > largest.Value {
				largest = &Product{Value: p, Factorizations: [][2]int{{x, y}}}
			} else if largest != nil && p == largest.Value {
				largest.Factorizations = append(largest.Factorizations, [2]int{x, y})
			}
		}
	}

	if smallest == nil || largest == nil {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	return *smallest, *largest, nil
}
