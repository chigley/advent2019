package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var memoryTests = []struct {
	initial []int
	final   []int
}{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

type testCase struct {
	inputs  []int
	outputs []int
}

var runTests = []struct {
	program []int
	tests   []testCase
}{
	{
		program: []int{3, 0, 4, 0, 99},
		tests: []testCase{
			{[]int{0}, []int{0}},
			{[]int{1}, []int{1}},
		},
	},
}

func TestMemory(t *testing.T) {
	for _, tt := range memoryTests {
		comp := New(tt.initial)
		if _, err := comp.Run(nil); err != nil {
			t.Error(err)
		}
		assert.Equal(t, tt.final, comp.memory)
	}
}

func TestRun(t *testing.T) {
	for _, tt := range runTests {
		comp := New(tt.program)
		for _, ttt := range tt.tests {
			outputs, err := comp.Run(ttt.inputs)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, ttt.outputs, outputs)
		}
	}
}
