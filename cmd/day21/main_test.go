package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day21 "github.com/chigley/advent2019/cmd/day21"
	"github.com/chigley/advent2019/intcode"
	"github.com/stretchr/testify/assert"
)

func TestDay19(t *testing.T) {
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

	part1, err := day21.Part1(comp)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 19349939, part1)
}
