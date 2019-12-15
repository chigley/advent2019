package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day15 "github.com/chigley/advent2019/cmd/day15"
	"github.com/stretchr/testify/assert"
)

func TestDay15(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	program, err := advent2019.ReadIntsLine(input)
	if err != nil {
		t.Fatal(err)
	}

	droid := day15.NewDroid(program)

	part1, err := droid.Part1()
	if err != nil {
		t.Error(err)
	}

	part2, err := droid.Part2()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 222, part1)
	assert.Equal(t, 394, part2)
}
