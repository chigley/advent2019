package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	galaxy, err := readGalaxy(input)
	if err != nil {
		t.Fatal(err)
	}

	// Part one
	assert.Equal(t, 42, orbits{"B": "COM", "C": "B", "D": "C", "E": "D", "F": "E", "G": "B", "H": "G", "I": "D", "J": "E", "K": "J", "L": "K"}.totalOrbits())
	assert.Equal(t, 292387, galaxy.totalOrbits())

	// Part two
	assert.Equal(t, 433, galaxy.orbitalTransfers("YOU", "SAN"))
	assert.Equal(t, 433, galaxy.orbitalTransfers("SAN", "YOU"))
}
