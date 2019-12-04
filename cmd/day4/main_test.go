package main

import (
	"testing"

	"github.com/chigley/advent2019"

	"github.com/stretchr/testify/assert"
)

func TestMeetsCriteria(t *testing.T) {
	assert.True(t, meetsCriteria(111111))
	assert.False(t, meetsCriteria(223450))
	assert.False(t, meetsCriteria(123789))

	assert.Equal(t, 1650, part1(advent2019.Range{From: 178416, To: 676461}))
}
