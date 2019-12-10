package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
)

func main() {
	input, err := advent2019.ReadStrings(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(Asteroids(input)))
}

func Part1(asteroids []advent2019.Point) int {
	mostAsteroids := -1

	for _, a1 := range asteroids {
		uniqueDirs := make(map[advent2019.Direction]struct{})
		for _, a2 := range asteroids {
			if a1 == a2 {
				continue
			}
			uniqueDirs[a1.Direction(a2)] = struct{}{}
		}
		mostAsteroids = advent2019.Max(mostAsteroids, len(uniqueDirs))
	}

	return mostAsteroids
}

func Asteroids(input []string) (ret []advent2019.Point) {
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				ret = append(ret, advent2019.Point{X: x, Y: y})
			}
		}
	}
	return ret
}
