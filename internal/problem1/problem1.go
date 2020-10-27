package problem1

func SumMultiplesOfThreeAndFive(rng int) int {
	var sum int
	for i := 1; i <= rng; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
