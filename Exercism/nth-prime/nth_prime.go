package prime

const Limit = 200000

var Primes map[int]bool

func Nth(n int) (int, bool) {
	InitPrimes(Limit)
	for i := 2; i <= Limit; i++ {
		for j := i + i; j <= Limit; j += i {
			Primes[j] = false
		}
	}

	var index int
	for k := 1; k <= Limit; k++ {
		if Primes[k] {
			index++
			if index == n {
				return k, true
			}
		}
	}
	return 0, false
}

func InitPrimes(limit int) {
	Primes = make(map[int]bool)
	for i := 2; i <= limit; i++ {
		Primes[i] = true
	}
}
