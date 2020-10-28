package problem3

func LargestPrimeFactor(n int) int {
	primes := getPrimeNumbers(n / 2)

	for i := len(primes); i > 0; i-- {
		p := primes[i-1]
		if n%p == 0 && p != 1 {
			return p
		}
	}

	return n
}

func getPrimeNumbers(n int) []int {
	primes := []int{1}
	for i := 2; i <= n; i++ {
		isPrime := true
		for _, p := range primes[1:] {
			if i%p == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
