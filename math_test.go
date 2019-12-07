package advent2019_test

import (
	"testing"

	"github.com/chigley/advent2019"
	"github.com/stretchr/testify/assert"
)

var permutationTests = []struct {
	input        []int
	outputLength int
}{
	{nil, 0},
	{[]int{}, 0},
	{[]int{0}, 1},
	{[]int{0, 1}, 2},
	{[]int{0, 1, 2}, 6},
	{[]int{0, 1, 2, 3}, 24},
}

func TestPermutations(t *testing.T) {
	for _, tt := range permutationTests {
		var total int
		c := advent2019.Permutations(tt.input)
		if c != nil {
			for range advent2019.Permutations(tt.input) {
				total++
			}
		}
		assert.Equal(t, tt.outputLength, total)
	}
}
