package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
)

func main() {
	masses, err := advent2019.ReadInts(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(masses))
}

func part1(masses []int) int {
	var totalFuel int
	for _, m := range masses {
		totalFuel += fuel(m)
	}
	return totalFuel
}

func fuel(mass int) int {
	return mass/3 - 2
}
