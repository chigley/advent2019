package main_test

import (
	"os"
	"testing"

	day8 "github.com/chigley/advent2019/cmd/day8"
	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	img, err := day8.ReadImage(input, 25, 6)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2159, day8.Part1(img))
}
