package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day23 "github.com/chigley/advent2019/cmd/day23"
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

	assert.Equal(t, 16250, day23.Part1(program))
}
