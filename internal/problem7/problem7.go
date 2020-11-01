package problem7

func NthPrime(n int) int {
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes[len(primes)-1]
}
