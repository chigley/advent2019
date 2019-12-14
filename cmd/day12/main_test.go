package main_test

import (
	"os"
	"testing"

	day12 "github.com/chigley/advent2019/cmd/day12"
	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	moons day12.Moons
	steps int

	part1, part2 int
}{
	{
		moons: day12.Moons{
			X: day12.Axis{{Pos: -1}, {Pos: 2}, {Pos: 4}, {Pos: 3}},
			Y: day12.Axis{{Pos: 0}, {Pos: -10}, {Pos: -8}, {Pos: 5}},
			Z: day12.Axis{{Pos: 2}, {Pos: -7}, {Pos: 8}, {Pos: -1}},
		},
		steps: 10,
		part1: 179,
		part2: 2772,
	},
	{
		moons: day12.Moons{
			X: day12.Axis{{Pos: -8}, {Pos: 5}, {Pos: 2}, {Pos: 9}},
			Y: day12.Axis{{Pos: -10}, {Pos: 5}, {Pos: -7}, {Pos: -8}},
			Z: day12.Axis{{Pos: 0}, {Pos: 10}, {Pos: 3}, {Pos: -3}},
		},
		steps: 100,
		part1: 1940,
		part2: 4686774924,
	},
}

func TestDay12(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	moons, err := day12.ReadMoons(input)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 7138, day12.Part1(moons.Clone(), 1000))
	assert.Equal(t, 572087463375796, day12.Part2(moons.Clone()))
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		assert.Equal(t, tt.part1, day12.Part1(tt.moons.Clone(), tt.steps))
		assert.Equal(t, tt.part2, day12.Part2(tt.moons))
	}
}
