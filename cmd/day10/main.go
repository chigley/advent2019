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

	_, part1 := Part1(Asteroids(input))
	fmt.Println(part1)
}

func Part1(asteroids []advent2019.Point) (advent2019.Point, int) {
	var p advent2019.Point
	mostAsteroids := -1

	for _, a1 := range asteroids {
		uniqueDirs := make(map[advent2019.Direction]struct{})
		for _, a2 := range asteroids {
			if a1 == a2 {
				continue
			}
			uniqueDirs[a1.Direction(a2)] = struct{}{}
		}
		if numAsteroids := len(uniqueDirs); numAsteroids > mostAsteroids {
			p = a1
			mostAsteroids = numAsteroids
		}
		mostAsteroids = advent2019.Max(mostAsteroids, len(uniqueDirs))
	}

	return p, mostAsteroids
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
