package problem3

import "errors"

func LargestPrimeFactor(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Range end must be greater than 0")
	}

	c := make(chan int)
	primes := make([]int, 0)

	go getPrimeNumbers(n, c)
	for p := range c {
		if n == p {
			c <- 1
			break
		}

		c <- 0

		primes = append(primes, p)
		n = tryKnownPrimes(n, primes)
	}

	return n, nil
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
			c <- i
			if <-c == 1 {
				close(c)
				return
			}
		}
	}
}
