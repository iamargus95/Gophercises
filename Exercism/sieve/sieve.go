package sieve

func Sieve(n int) []int {
	primes := make([]int, 0)
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if sieve[i] {
			continue
		}
		primes = append(primes, i)
		for j := i; j <= n; j += i {
			sieve[j] = true
		}
	}
	return primes
}
