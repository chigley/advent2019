package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	initial []int
	final   []int
}{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func TestRun(t *testing.T) {
	for _, tt := range tests {
		comp := New(tt.initial)
		if err := comp.Run(); err != nil {
			t.Error(err)
		}
		assert.Equal(t, tt.final, comp.memory)
	}
}
