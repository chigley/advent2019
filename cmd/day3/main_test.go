package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	wires [][]string
	dist  distance
	steps numSteps
}{
	{
		wires: [][]string{
			{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
		},
		dist:  159,
		steps: 610,
	},
	{
		wires: [][]string{
			{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
		},
		dist:  135,
		steps: 410,
	},
}

func TestOptimalPoint(t *testing.T) {
	for _, tt := range tests {
		dist, steps, err := optimalPoint(tt.wires)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, tt.dist, dist)
		assert.Equal(t, tt.steps, steps)
	}
}
