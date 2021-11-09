// Package summultiples sums multiples.
package summultiples

// SumMultiples sums multiples.
func SumMultiples(limit int, divisors ...int) int {
	multiples := map[int]struct{}{}
	find := func(orig int) {
		v := orig
		for v < limit {
			multiples[v] = struct{}{}
			v += orig
		}
	}

	for _, v := range divisors {
		if v == 0 {
			continue
		}
		find(v)
	}

	var result int
	for k := range multiples {
		result += k
	}

	return result
}
