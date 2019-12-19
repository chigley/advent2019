package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day19 "github.com/chigley/advent2019/cmd/day19"
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

	part1, err := day19.Part1(program)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 186, part1)
}
