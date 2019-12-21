package main_test

import (
	"os"
	"testing"

	day18 "github.com/chigley/advent2019/cmd/day18"
	"github.com/stretchr/testify/assert"
)

type part1Test struct {
	path   string
	output int
}

var part1Tests = []part1Test{
	{"testdata/example0", 8},
	{"testdata/example1", 86},
	{"testdata/example2", 132},
	{"testdata/example3", 136},
	{"testdata/example4", 81},
	{"testdata/input", 5182},
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		tt.run(t)
	}
}

func (tt part1Test) run(t *testing.T) {
	input, err := os.Open(tt.path)
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	maze, err := day18.NewMaze(input)
	if err != nil {
		t.Error(err)
	}

	output, err := maze.Part1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, tt.output, output)
}
