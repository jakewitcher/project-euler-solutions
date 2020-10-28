package main

import (
	"fmt"

	"github.com/jakewitcher/project-euler-solutions/internal/problem1"
	"github.com/jakewitcher/project-euler-solutions/internal/problem2"
	"github.com/jakewitcher/project-euler-solutions/internal/problem3"
)

func main() {
	FindLargestPrimeFactor()
	SumMultiplesOfThreeAndFiveRange1000()
	SumEvenFibonacciNumbersMaxFourMillion()
}

func SumMultiplesOfThreeAndFiveRange1000() {
	sum := problem1.SumMultiplesOfThreeAndFive(1000)
	fmt.Printf("Sum of multiples of 3 and 5, ranging from 1 to 1000 : %v\n", sum)
}

func SumEvenFibonacciNumbersMaxFourMillion() {
	sum := problem2.SumEvenFibonacciNumbers(4_000_000)
	fmt.Printf("Sum of even fibonacci numbers, from 1 to 4 million : %v\n", sum)
}

func FindLargestPrimeFactor() {
	prime := problem3.LargestPrimeFactor(600_851_475_143)
	fmt.Printf("Largest prime factor for 600,851,475,143 : %v\n", prime)
}
