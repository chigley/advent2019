package main

import (
	"fmt"

	"github.com/chigley/advent2019/vector"
)

type pos vector.XY

type direction int

const (
	_ direction = iota
	north
	south
	west
	east
)

var directions = []direction{north, south, west, east}

func (p pos) next(d direction) (pos, error) {
	switch d {
	case north:
		return pos{X: p.X, Y: p.Y + 1}, nil
	case south:
		return pos{X: p.X, Y: p.Y - 1}, nil
	case east:
		return pos{X: p.X + 1, Y: p.Y}, nil
	case west:
		return pos{X: p.X - 1, Y: p.Y}, nil
	default:
		return pos{}, fmt.Errorf("unrecognised direction %d", d)
	}
}

func (d direction) opposite() (direction, error) {
	switch d {
	case north:
		return south, nil
	case south:
		return north, nil
	case east:
		return west, nil
	case west:
		return east, nil
	default:
		return 0, fmt.Errorf("unrecognised direction %d", d)
	}
}

func (d direction) String() string {
	switch d {
	case north:
		return "north"
	case south:
		return "south"
	case east:
		return "east"
	case west:
		return "west"
	default:
		return fmt.Sprintf("UnknownDirection(%d)", d)
	}
}
