package prime

func Factors(n int64) []int64 {
	factors := []int64{}

	var c int64 = 3
	for n > 1 {

		for n%2 == 0 {
			factors = append(factors, 2)
			n /= 2
		}

		if n%c == 0 {
			factors = append(factors, c)
			n /= c
			continue
		}

		c += 2
	}

	return factors
}
