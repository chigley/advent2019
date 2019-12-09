package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	// Using position mode, consider whether the input is equal to 8; output 1
	// (if it is) or 0 (if it is not).
	{
		program: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		tests: []testCase{
			{[]int{7}, []int{0}},
			{[]int{8}, []int{1}},
			{[]int{9}, []int{0}},
		},
	},

	// Using position mode, consider whether the input is less than 8; output 1
	// (if it is) or 0 (if it is not).
	{
		program: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		tests: []testCase{
			{[]int{7}, []int{1}},
			{[]int{8}, []int{0}},
			{[]int{9}, []int{0}},
		},
	},

	// Using immediate mode, consider whether the input is equal to 8; output 1
	// (if it is) or 0 (if it is not).
	{
		program: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		tests: []testCase{
			{[]int{7}, []int{0}},
			{[]int{8}, []int{1}},
			{[]int{9}, []int{0}},
		},
	},

	// Using immediate mode, consider whether the input is less than 8; output 1
	// (if it is) or 0 (if it is not).
	{
		program: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		tests: []testCase{
			{[]int{7}, []int{1}},
			{[]int{8}, []int{0}},
			{[]int{9}, []int{0}},
		},
	},

	// Take an input, then output 0 if the input was zero or 1 if the input was
	// non-zero.
	{
		program: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		tests: []testCase{
			{[]int{-1}, []int{1}},
			{[]int{0}, []int{0}},
			{[]int{1}, []int{1}},
		},
	},
	{
		program: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		tests: []testCase{
			{[]int{-1}, []int{1}},
			{[]int{0}, []int{0}},
			{[]int{1}, []int{1}},
		},
	},

	// Output 999 if the input value is below 8, output 1000 if the input value
	// is equal to 8, or output 1001 if the input value is greater than 8.
	{
		program: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20,
			1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105,
			1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46,
			98, 99},
		tests: []testCase{
			{[]int{7}, []int{999}},
			{[]int{8}, []int{1000}},
			{[]int{9}, []int{1001}},
		},
	},

	// Takes no input and produces a copy of itself as output.
	{
		program: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101,
			1006, 101, 0, 99},
		tests: []testCase{
			{[]int{7}, []int{109, 1, 204, -1, 1001, 100, 1, 100,
				1008, 100, 16, 101, 1006, 101, 0, 99}},
		},
	},
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
