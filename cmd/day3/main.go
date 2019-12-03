package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/chigley/advent2019"
)

var (
	dirX = map[string]int{"R": 1, "U": 0, "L": -1, "D": 0}
	dirY = map[string]int{"R": 0, "U": 1, "L": 0, "D": -1}
)

type point struct {
	x, y int
}

type (
	distance int
	numSteps int

	wireID int

	visitedPoints map[point]wireSet
	wireSet       map[wireID]numSteps
)

func main() {
	wires, err := advent2019.ReadCSV(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, part2, err := optimalPoint(wires)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func optimalPoint(wires [][]string) (distance, numSteps, error) {
	visited := make(visitedPoints)
	for i, path := range wires {
		if err := visited.walkWire(wireID(i), path); err != nil {
			return 0, 0, err
		}
	}
	return visited.optimalPoint()
}

func (v visitedPoints) optimalPoint() (distance, numSteps, error) {
	var (
		origin point

		maxVisited  int
		minDistance distance
		minSteps    numSteps
	)

	for point, wires := range v {
		if len(wires) < maxVisited {
			// We already found a point that more wires visited than this one.
			continue
		}

		dist := origin.distance(point)

		var totalSteps numSteps
		for _, steps := range wires {
			totalSteps += steps
		}

		if len(wires) > maxVisited {
			// The current point has been visited by more wires than all
			// previous points. This is our new favourite, no matter how far
			// away it was.
			maxVisited = len(wires)
			minDistance = dist
			minSteps = totalSteps
			continue
		}

		if dist < minDistance {
			minDistance = dist
		}
		if totalSteps < minSteps {
			minSteps = totalSteps
		}
	}

	if maxVisited < 2 {
		return 0, 0, errors.New("expected to find a point visited by at least two wires")
	}
	return minDistance, minSteps, nil
}

func (v visitedPoints) walkWire(id wireID, path []string) error {
	var (
		pos   point
		steps numSteps
	)
	for _, move := range path {
		if len(move) < 2 {
			return fmt.Errorf("move %s is too short", move)
		}

		dir, distanceStr := move[:1], move[1:]
		x, ok := dirX[dir]
		if !ok {
			return fmt.Errorf("unrecognised direction %s", dir)
		}
		y, ok := dirY[dir]
		if !ok {
			return fmt.Errorf("unrecognised direction %s", dir)
		}

		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			return err
		}

		for i := 0; i < distance; i++ {
			pos.x += x
			pos.y += y
			steps++

			wires, ok := v[pos]
			if ok {
				// A wire - perhaps ourseves - already visited here. Extend
				// existing map if it wasn't us.
				if _, ok2 := wires[id]; !ok2 {
					wires[id] = steps
				}
			} else {
				// We're the first wire to visit this point. Create new map.
				v[pos] = wireSet{id: steps}
			}
		}
	}
	return nil
}

func (p1 *point) distance(p2 point) distance {
	return distance(advent2019.Abs(p1.x-p2.x) + advent2019.Abs(p1.y-p2.y))
}
