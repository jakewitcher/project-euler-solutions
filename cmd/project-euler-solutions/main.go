package main

import (
	"fmt"

	"github.com/jakewitcher/project-euler-solutions/internal/problem-1/problem1"
)

func main() {

}

func FindMultiplesOfThreeAndFiveRange1000() {
	multiples := problem1.FindMultiplesOfThreeAndFive(1000)
	fmt.Println("Multiples of 3 and 5, from 1 to 1000")
	for _, n := range multiples {
		fmt.Println(n)
	}
}
