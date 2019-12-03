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

	wireID int

	visitedPoints map[point]wireSet
	wireSet       map[wireID]struct{}
)

func main() {
	wires, err := advent2019.ReadCSV(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := optimalPoint(wires)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
}

func optimalPoint(wires [][]string) (distance, error) {
	visited := make(visitedPoints)
	for i, path := range wires {
		if err := visited.walkWire(wireID(i), path); err != nil {
			return 0, err
		}
	}
	return visited.optimalPoint()
}

func (v visitedPoints) optimalPoint() (distance, error) {
	var (
		origin point

		maxVisited  int
		minDistance distance
	)

	for point, wires := range v {
		if len(wires) < maxVisited {
			// We already found a point that more wires visited than this one.
			continue
		}

		dist := origin.distance(point)

		if len(wires) > maxVisited {
			// The current point has been visited by more wires than all
			// previous points. This is our new favourite, no matter how far
			// away it was.
			maxVisited = len(wires)
			minDistance = dist
			continue
		}

		if dist < minDistance {
			minDistance = dist
		}
	}

	if maxVisited < 2 {
		return 0, errors.New("expected to find a point visited by at least two wires")
	}
	return minDistance, nil
}

func (v visitedPoints) walkWire(id wireID, path []string) error {
	var pos point
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

			wires, ok := v[pos]
			if ok {
				// A wire has already visited this point. Add ourselves to the
				// set.
				wires[id] = struct{}{}
			} else {
				// We're the first wire to visit this point.
				v[pos] = wireSet{id: struct{}{}}
			}
		}
	}
	return nil
}

func (p1 *point) distance(p2 point) distance {
	return distance(advent2019.Abs(p1.x-p2.x) + advent2019.Abs(p1.y-p2.y))
}
