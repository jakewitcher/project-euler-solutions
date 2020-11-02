package problem10

func SummationOfPrimes(max int) int {
	nums := make([]int, 0, max/3)
	nums = append(nums, 2, 3, 5, 7)

	primes := make([]int, 0, max/4)
	primes = append(primes, 2, 3, 5, 7)

	for i := 11; i < max; i += 2 {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			continue
		}
		nums = append(nums, i)
	}

	for _, n := range nums {
		isPrime := true
		for _, p := range primes {
			if n%p == 0 {
				isPrime = false
				break
			}
		}

		if isPrime && n < max {
			primes = append(primes, n)
		}
	}

	sum := 1
	for _, p := range primes {
		sum += p
	}

	return sum
}
