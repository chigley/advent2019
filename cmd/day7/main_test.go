package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day7 "github.com/chigley/advent2019/cmd/day7"
	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	program []int
	output  int
}{
	{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
	{[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, 54321},
	{[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, 65210},
}

func TestDay7(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	program, err := advent2019.ReadIntsLine(input)
	if err != nil {
		t.Fatal(err)
	}

	part1, err := day7.Part1(program)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 880726, part1)
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		output, err := day7.Part1(tt.program)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tt.output, output)
	}
}
