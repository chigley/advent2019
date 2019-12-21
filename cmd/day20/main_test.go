package main_test

import (
	"os"
	"testing"

	day20 "github.com/chigley/advent2019/cmd/day20"
	"github.com/stretchr/testify/assert"
)

type test struct {
	path   string
	output int
}

var part1Tests = []test{
	{"testdata/example0", 23},
	{"testdata/example1", 58},
	{"testdata/input", 442},
}

var part2Tests = []test{
	{"testdata/example0", 26},
	{"testdata/example2", 396},
	{"testdata/input", 5208},
}

func TestDay20(t *testing.T) {
	for _, tt := range part1Tests {
		tt.run1(t)
	}
	for _, tt := range part2Tests {
		tt.run2(t)
	}
}

func (tt test) run1(t *testing.T) {
	input, err := os.Open(tt.path)
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	maze, err := day20.NewMaze(input)
	if err != nil {
		t.Error(err)
	}

	output, err := maze.Part1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, tt.output, output)
}

func (tt test) run2(t *testing.T) {
	input, err := os.Open(tt.path)
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	maze, err := day20.NewMaze(input)
	if err != nil {
		t.Error(err)
	}

	output, err := maze.Part2()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, tt.output, output)
}
