package problem4

import (
	"errors"
	"math"
	"strconv"
)

func LargestPalindromeProduct(digits float64) (int, error) {
	if digits < 1 {
		return 0, errors.New("number of digits must be greater than 0")
	}

	rangeBegin := int(math.Pow(10, digits-1))
	rangeEnd := int(math.Pow(10, digits))

	palindromes := make([]int, 0)

	for i := rangeBegin; i < rangeEnd; i++ {
		for j := i; j < rangeEnd; j++ {
			product := j * i

			if isPalindrome(product) {
				palindromes = append(palindromes, product)
			}
		}
	}

	var largest int
	for _, p := range palindromes {
		if p > largest {
			largest = p
		}
	}

	return largest, nil
}

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	return str == reverse(str)

}

func reverse(s string) string {
	reverse := make([]rune, len(s))
	last := len(s) - 1
	for i, r := range s {
		reverse[last-i] = r
	}

	return string(reverse)
}
