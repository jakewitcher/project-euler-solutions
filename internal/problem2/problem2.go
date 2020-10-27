package problem2

func SumEvenFibonacciNumbers(max int) int {
	var sum int

	for prev, curr := 1, 2; curr <= max; {
		if curr%2 == 0 {
			sum += curr
		}

		prev = curr
		curr += prev
	}

	return sum
}
