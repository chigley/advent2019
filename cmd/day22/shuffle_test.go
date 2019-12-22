package main_test

import (
	"testing"

	day22 "github.com/chigley/advent2019/cmd/day22"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	// stack
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, day22.Stack{}.Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, []int{8, 7, 6, 5, 4, 3, 2, 1, 0}, day22.Stack{}.Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))

	// cut
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, day22.Cut{N: 0}.Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, day22.Cut{N: 3}.Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, day22.Cut{N: -4}.Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))

	// increment
	assert.Equal(t, []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, day22.Increment{N: 3}.Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
