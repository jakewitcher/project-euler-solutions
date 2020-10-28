package problem2

import "errors"

func SumEvenFibonacciNumbers(max int) (int, error) {
	if max < 1 {
		return 0, errors.New("Range end must be greater than 0")
	}

	var sum int

	for prev, curr := 1, 2; curr <= max; {
		if curr%2 == 0 {
			sum += curr
		}

		prev = curr
		curr += prev
	}

	return sum, nil
}
