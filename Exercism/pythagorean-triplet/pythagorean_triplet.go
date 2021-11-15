package pythagorean

import "math"

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	res := []Triplet{}

	for i := min; i <= max-2; i++ {

		for j := i + 1; j <= max-1; j++ {

			for k := j + 1; k <= max; k++ {

				if math.Pow(float64(i), 2)+math.Pow(float64(j), 2) == math.Pow(float64(k), 2) {

					res = append(res, Triplet{i, j, k})
				}
			}
		}
	}
	return res
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.
func Sum(p int) []Triplet {
	res := []Triplet{}

	for i := 1; i <= p/3; i++ {

		for j := i + 1; j <= p-i; j++ {

			k := p - i - j

			if k < j {

				continue
			}

			if math.Pow(float64(i), 2)+math.Pow(float64(j), 2) == math.Pow(float64(k), 2) {

				res = append(res, Triplet{i, j, k})
			}
		}
	}

	return res
}
