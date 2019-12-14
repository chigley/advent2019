package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day13 "github.com/chigley/advent2019/cmd/day13"
	"github.com/stretchr/testify/assert"
)

func TestDay13(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	program, err := advent2019.ReadIntsLine(input)
	if err != nil {
		t.Fatal(err)
	}

	part1, err := day13.Part1(program)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 369, part1)
}
