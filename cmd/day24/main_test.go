package main_test

import (
	"os"
	"testing"

	day24 "github.com/chigley/advent2019/cmd/day24"
	"github.com/stretchr/testify/assert"
)

func TestDay24(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	g, err := day24.ReadGrid(input)
	if err != nil {
		t.Error(err)
	}

	t.Log(g)

	assert.Equal(t, uint64(32509983), day24.Part1(g))
}
