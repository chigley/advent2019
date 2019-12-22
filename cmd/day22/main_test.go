package main_test

import (
	"os"
	"testing"

	day22 "github.com/chigley/advent2019/cmd/day22"
	"github.com/stretchr/testify/assert"
)

func TestDay22(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	order, err := day22.ReadOrder(input)
	if err != nil {
		t.Error(err)
	}

	part1, err := order.Part1()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 4775, part1)
}
