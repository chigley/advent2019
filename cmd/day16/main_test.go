package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"

	day16 "github.com/chigley/advent2019/cmd/day16"
	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	input  string
	output string
}{
	{input: "80871224585914546619083218645595", output: "24176176"},
	{input: "19617804207202209144916044189917", output: "73745418"},
	{input: "69317163492948606335995924319873", output: "52432133"},
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
	assert.Equal(t, "15841929", part1)
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
