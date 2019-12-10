package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"github.com/chigley/advent2019"
)

func main() {
	input, err := advent2019.ReadStrings(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	asteroids := Asteroids(input)

	laserPos, part1 := Part1(asteroids)

	part2, err := Part2(laserPos, asteroids)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2.X*100 + part2.Y)
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
			dir, _ := a1.Direction(a2)
			uniqueDirs[dir] = struct{}{}
		}
		if numAsteroids := len(uniqueDirs); numAsteroids > mostAsteroids {
			p = a1
			mostAsteroids = numAsteroids
		}
		mostAsteroids = advent2019.Max(mostAsteroids, len(uniqueDirs))
	}

	return p, mostAsteroids
}

func Part2(laserPos advent2019.Point, input []advent2019.Point) (advent2019.Point, error) {
	hits := HitOrder(laserPos, input)
	if len(hits) < 200 {
		return advent2019.Point{}, errors.New("not enough asteroids")
	}
	return hits[199], nil
}

func HitOrder(laserPos advent2019.Point, input []advent2019.Point) []advent2019.Point {
	type asteroid struct {
		pos  advent2019.Point
		dist int
	}

	// Group asteroids by their angle
	asteroidsAtAngle := make(map[float64][]asteroid)
	for _, a := range input {
		if a == laserPos {
			continue
		}

		dir, dist := laserPos.Direction(a)

		angle := math.Atan2(float64(dir.DX), float64(-dir.DY))
		if angle < 0 {
			angle += 2 * math.Pi
		}

		cur := asteroidsAtAngle[angle]
		asteroidsAtAngle[angle] = append(cur, asteroid{
			pos:  a,
			dist: dist,
		})
	}

	// 1. Build a sorted slice, angles, of the keys of asteroidsAtAngle
	// 2. Sort each asteroid slice by ascending distance
	var angles sort.Float64Slice
	for angle, asteroids := range asteroidsAtAngle {
		angles = append(angles, angle)
		sort.Slice(asteroids, func(i, j int) bool { return asteroids[i].dist < asteroids[j].dist })
	}
	angles.Sort()

	// Loop through the angles in order, popping the nearest asteroid from the
	// front of the slice each time
	hits := make([]advent2019.Point, 0, len(input)-1)
	for {
		for _, ang := range angles {
			asteroids := asteroidsAtAngle[ang]
			if len(asteroids) == 0 {
				continue
			}

			hit, rest := asteroids[0], asteroids[1:]
			hits = append(hits, hit.pos)
			asteroidsAtAngle[ang] = rest
			if len(hits) == len(input)-1 {
				return hits
			}
		}
	}
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
