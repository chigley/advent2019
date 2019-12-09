package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day9 "github.com/chigley/advent2019/cmd/day9"
	"github.com/chigley/advent2019/intcode"
	"github.com/stretchr/testify/assert"
)

func TestDay9(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	program, err := advent2019.ReadIntsLine(input)
	if err != nil {
		t.Fatal(err)
	}

	comp := intcode.New(program)

	part1, err := day9.Part1(comp)
	if err != nil {
		t.Error(err)
	}

	part2, err := day9.Part2(comp)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2682107844, part1)
	assert.Equal(t, 34738, part2)
}
