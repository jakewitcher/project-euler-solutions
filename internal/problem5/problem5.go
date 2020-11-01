package problem5

import "math"

type Factors = map[int]int

func SmallestMultiple(end int) int {
	allFactors := make(Factors)

	for i := 2; i <= end; i++ {
		appendFactors(allFactors, i)

	}

	product := 1

	for factor, amt := range allFactors {
		product *= int(math.Pow(float64(factor), float64(amt)))
	}

	return product
}

func appendFactors(allFactors Factors, n int) {
	factors := getFactors(make(Factors), allFactors, n)

	for factor, amt := range factors {
		if n := allFactors[factor]; n < amt {
			allFactors[factor] = amt
		}
	}
}

func getFactors(factors Factors, allFactors Factors, n int) Factors {
	for p := range allFactors {
		if n%p == 0 && n != 1 {
			factors[p]++
			return getFactors(factors, allFactors, n/p)
		}
	}

	if len(factors) == 0 {
		factors[n]++
	}

	return factors
}
