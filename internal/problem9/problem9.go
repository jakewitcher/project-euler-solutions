package problem9

func SpecialPythagoreanTriplet() int {
	sum := 1000
	for a := 1; a <= sum; a++ {
		for b := 1; b <= sum; b++ {
			c := sum - a - b
			if isPythagoreanTriplet(a, b, c) {
				return a * b * c
			}
		}
	}

	return 0
}

func isPythagoreanTriplet(a, b, c int) bool {
	return a < b && b < c && a*a+b*b == c*c
}
