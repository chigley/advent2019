package main

import (
	"testing"

	"github.com/chigley/advent2019"

	"github.com/stretchr/testify/assert"
)

func TestMeetsCriteria(t *testing.T) {
	assert.True(t, meetsPart1Criteria("111111"))
	assert.False(t, meetsPart1Criteria("223450"))
	assert.False(t, meetsPart1Criteria("123789"))

	assert.True(t, meetsPart2Criteria("112233"))
	assert.False(t, meetsPart2Criteria("123444"))
	assert.True(t, meetsPart2Criteria("111122"))

	part1, part2 := solve(advent2019.Range{From: 178416, To: 676461})
	assert.Equal(t, 1650, part1)
	assert.Equal(t, 1129, part2)
}
