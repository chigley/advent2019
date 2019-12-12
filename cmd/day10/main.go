package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
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

func Part1(asteroids []vector.XY) (vector.XY, int) {
	var p vector.XY
	mostAsteroids := -1

	for _, a1 := range asteroids {
		uniqueDirs := make(map[vector.XY]struct{})
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

func Part2(laserPos vector.XY, input []vector.XY) (vector.XY, error) {
	hits := HitOrder(laserPos, input)
	if len(hits) < 200 {
		return vector.XY{}, errors.New("not enough asteroids")
	}
	return hits[199], nil
}

func HitOrder(laserPos vector.XY, input []vector.XY) []vector.XY {
	type asteroid struct {
		pos  vector.XY
		dist int
	}

	// Group asteroids by their angle
	asteroidsAtAngle := make(map[float64][]asteroid)
	for _, a := range input {
		if a == laserPos {
			continue
		}

		dir, dist := laserPos.Direction(a)

		angle := math.Atan2(float64(dir.X), float64(-dir.Y))
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
	hits := make([]vector.XY, 0, len(input)-1)
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

func Asteroids(input []string) (ret []vector.XY) {
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				ret = append(ret, vector.XY{X: x, Y: y})
			}
		}
	}
	return ret
}
