package problem6

func SumSquareDifference(end int) int {
	sum := 0
	sumOfSquares := 0
	for i := 1; i <= end; i++ {
		sum += i
		sumOfSquares += i * i
	}

	return sum*sum - sumOfSquares
}
