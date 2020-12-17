package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	cells := parseInput(testInput)
	assert.Equal(t, 5, len(cells))
}

func TestIterate(t *testing.T) {
	cells := parseInput(testInput)
	for i := 0; i < 6; i++ {
		cells = iterate(cells)
	}
	assert.Equal(t, 848, len(cells))
}

var testInput = `
.#.
..#
###
`
