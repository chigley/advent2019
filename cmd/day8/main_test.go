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

	if _, err := day8.Part2(img); err != nil {
		t.Error(err)
	}
}

func TestRender(t *testing.T) {
	layers := day8.SpaceImage{
		{[]day8.Pixel{0, 2}, []day8.Pixel{2, 2}},
		{[]day8.Pixel{1, 1}, []day8.Pixel{2, 2}},
		{[]day8.Pixel{2, 2}, []day8.Pixel{1, 2}},
		{[]day8.Pixel{0, 0}, []day8.Pixel{0, 0}},
	}

	image, err := layers.Render()
	if err != nil {
		t.Error(err)
	}

	expected := day8.Layer{
		{0, 1},
		{1, 0},
	}
	assert.Equal(t, expected, image)
}
