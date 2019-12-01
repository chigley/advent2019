package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
)

type fuelFunc func(mass int) (fuel int)

func main() {
	masses, err := advent2019.ReadInts(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(masses))
	fmt.Println(part2(masses))
}

func totalFuel(masses []int, f fuelFunc) int {
	var totalFuel int
	for _, m := range masses {
		totalFuel += f(m)
	}
	return totalFuel
}

func part1(masses []int) int {
	return totalFuel(masses, fuel)
}

func part2(masses []int) int {
	return totalFuel(masses, fuelRecursive)
}

func fuel(mass int) int {
	return mass/3 - 2
}

func fuelRecursive(mass int) int {
	totalFuel := fuel(mass)
	for extra := fuel(totalFuel); extra > 0; extra = fuel(extra) {
		totalFuel += extra
	}
	return totalFuel
}
