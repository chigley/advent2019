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

var part2Tests = []struct {
	program []int
	output  int
}{
	{
		[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4,
			27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		139629729,
	},
	{
		[]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55,
			1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008,
			54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56,
			1005, 56, 6, 99, 0, 0, 0, 0, 10}, 18216,
	},
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

	part2, err := day7.Part2(program)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 880726, part1)
	assert.Equal(t, 4931744, part2)

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

func TestPart2(t *testing.T) {
	for _, tt := range part2Tests {
		output, err := day7.Part2(tt.program)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tt.output, output)
	}
}
