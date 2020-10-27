package main

import (
	"fmt"

	"github.com/jakewitcher/project-euler-solutions/internal/problem1"
	"github.com/jakewitcher/project-euler-solutions/internal/problem2"
)

func main() {
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
