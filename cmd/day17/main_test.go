package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day17 "github.com/chigley/advent2019/cmd/day17"
	"github.com/stretchr/testify/assert"
)

const (
	part1In = `..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..`
	part1Out = 76
)

func TestDay17(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	program, err := advent2019.ReadIntsLine(input)
	if err != nil {
		t.Fatal(err)
	}

	view, err := day17.ReadView(program)
	if err != nil {
		t.Error(err)
	}

	part1, err := day17.Part1(view)
	if err != nil {
		t.Error(err)
	}

	part2, err := day17.Part2(program)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 5680, part1)
	assert.Equal(t, 895965, part2)
}

func TestPart1(t *testing.T) {
	part1, err := day17.Part1(part1In)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, part1Out, part1)
}
