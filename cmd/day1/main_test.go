package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel(t *testing.T) {
	assert.Equal(t, 2, fuel(12))
	assert.Equal(t, 2, fuel(14))
	assert.Equal(t, 654, fuel(1969))
	assert.Equal(t, 33583, fuel(100756))
}
