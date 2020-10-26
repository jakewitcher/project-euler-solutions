package main

import (
	"fmt"

	"github.com/jakewitcher/project-euler-solutions/internal/problem1"
)

func main() {
	FindMultiplesOfThreeAndFiveRange1000()
}

func FindMultiplesOfThreeAndFiveRange1000() {
	multiples := problem1.FindMultiplesOfThreeAndFive(1000)
	fmt.Println("Multiples of 3 and 5, ranging from 1 to 1000")
	for _, n := range multiples {
		fmt.Println(n)
	}
}
