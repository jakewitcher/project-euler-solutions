package problem3

func LargestPrimeFactor(n int) int {
	c := make(chan int)
	primes := make([]int, 0)

	go getPrimeNumbers(n, c)
	for p := range c {
		if n == p {
			return n
		}

		primes = append(primes, p)
		n = tryKnownPrimes(n, primes)
	}

	return n
}

func tryKnownPrimes(n int, primes []int) int {
	for _, p := range primes {
		if n%p == 0 {
			return tryKnownPrimes(n/p, primes)
		}
	}

	return n
}

func getPrimeNumbers(n int, c chan int) {
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
			c <- i
		}
	}
}
