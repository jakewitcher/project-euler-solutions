package problem1

import "errors"

func SumMultiplesOfThreeAndFive(rng int) (int, error) {
	if rng < 1 {
		return 0, errors.New("Range end must be greater than 0")
	}

	var sum int
	for i := 1; i <= rng; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum, nil
}
