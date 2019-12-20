package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"

	day16 "github.com/chigley/advent2019/cmd/day16"
	"github.com/stretchr/testify/assert"
)

type test struct {
	input  string
	output string
}

var part1Tests = []test{
	{input: "80871224585914546619083218645595", output: "24176176"},
	{input: "19617804207202209144916044189917", output: "73745418"},
	{input: "69317163492948606335995924319873", output: "52432133"},
}

var part2Tests = []test{
	{input: "03036732577212944063491565474664", output: "84462026"},
	{input: "02935109699940807407585447034323", output: "78725270"},
	{input: "03081770884921959731165446850517", output: "53553731"},
}

func TestDay16(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	inputLines, err := advent2019.ReadStrings(input)
	if err != nil {
		t.Fatal(err)
	}

	digits, err := day16.ParseDigits(inputLines[0])
	if err != nil {
		t.Error(err)
	}

	part1, err := day16.Part1(digits)
	if err != nil {
		t.Error(err)
	}

	part2, err := day16.Part2(inputLines[0])
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "15841929", part1)
	assert.Equal(t, "39011547", part2)
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		digits, err := day16.ParseDigits(tt.input)
		if err != nil {
			t.Error(err)
		}

		output, err := day16.Part1(digits)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tt.output, output)
	}
}

func TestPart2(t *testing.T) {
	for _, tt := range part2Tests {
		output, err := day16.Part2(tt.input)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tt.output, output)
	}
}
