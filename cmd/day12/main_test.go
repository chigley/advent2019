package main_test

import (
	"os"
	"testing"

	day12 "github.com/chigley/advent2019/cmd/day12"
	"github.com/chigley/advent2019/vector"
	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	moons day12.Moons
	steps int

	output int
}{
	{
		moons: day12.Moons{
			day12.Moon{Pos: vector.XYZ{X: -1, Y: 0, Z: 2}},
			day12.Moon{Pos: vector.XYZ{X: 2, Y: -10, Z: -7}},
			day12.Moon{Pos: vector.XYZ{X: 4, Y: -8, Z: 8}},
			day12.Moon{Pos: vector.XYZ{X: 3, Y: 5, Z: -1}},
		},
		steps:  10,
		output: 179,
	},
	{
		moons: day12.Moons{
			day12.Moon{Pos: vector.XYZ{X: -8, Y: -10, Z: 0}},
			day12.Moon{Pos: vector.XYZ{X: 5, Y: 5, Z: 10}},
			day12.Moon{Pos: vector.XYZ{X: 2, Y: -7, Z: 3}},
			day12.Moon{Pos: vector.XYZ{X: 9, Y: -8, Z: -3}},
		},
		steps:  100,
		output: 1940,
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

	assert.Equal(t, 7138, day12.Part1(moons, 1000))
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		assert.Equal(t, tt.output, day12.Part1(tt.moons, tt.steps))
	}
}
