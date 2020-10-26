package problem1

func FindMultiplesOfThreeAndFive(rng int) []int {
	multiples := make([]int, 0)
	for i := 1; i <= rng; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples = append(multiples, i)
		}
	}

	return multiples
}
