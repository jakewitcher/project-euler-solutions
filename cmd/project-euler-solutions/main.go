package main

import (
	"fmt"

	"github.com/jakewitcher/project-euler-solutions/internal/problem1"
	"github.com/jakewitcher/project-euler-solutions/internal/problem2"
	"github.com/jakewitcher/project-euler-solutions/internal/problem3"
	"github.com/jakewitcher/project-euler-solutions/internal/problem4"
	"github.com/jakewitcher/project-euler-solutions/internal/problem5"
)

func main() {
	// SumMultiplesOfThreeAndFive()
	// SumEvenFibonacciNumbers()
	// FindLargestPrimeFactor()
	// FindLargestPalindrome()
	SmallesetMultiple()
}

func SumMultiplesOfThreeAndFive() {
	sum, _ := problem1.SumMultiplesOfThreeAndFive(1000)
	fmt.Printf("Sum of multiples of 3 and 5, ranging from 1 to 1000 : %v\n", sum)
}

func SumEvenFibonacciNumbers() {
	sum, _ := problem2.SumEvenFibonacciNumbers(4_000_000)
	fmt.Printf("Sum of even fibonacci numbers, from 1 to 4 million : %v\n", sum)
}

func FindLargestPrimeFactor() {
	prime, _ := problem3.LargestPrimeFactor(600_851_475_143)
	fmt.Printf("Largest prime factor for 600,851,475,143 : %v\n", prime)
}

func FindLargestPalindrome() {
	palindrome, _ := problem4.LargestPalindromeProduct(3)
	fmt.Printf("Largest palindrome product of three digit numbers : %v\n", palindrome)
}

func SmallesetMultiple() {
	multiple := problem5.SmallestMultiple(20)
	fmt.Printf("Smallest multiple of all numbers between 1 and 20 : %v\n", multiple)
}
