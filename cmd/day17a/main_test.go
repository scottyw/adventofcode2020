package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	cells := parseInput(c0z0, 0, 0, 0)
	assert.Equal(t, 5, len(cells))
}

func TestIterate(t *testing.T) {
	cycle0 := parseInput(c0z0, 0, 0, 0)
	cycle1 := iterate(cycle0)
	assert.Equal(t, c1, cycle1)
	cycle2 := iterate(cycle1)
	assert.Equal(t, c2, cycle2)
	cycle3 := iterate(cycle2)
	assert.Equal(t, c3, cycle3)
}

// Before any cycles:

var c0z0 = `
.#.
..#
###
`

// After 1 cycle: extra empty cells added at the top to keep coordinates the same across cycles

var c1zn1 = `
#..
..#
.#.
`

var c1z0 = `
#.#
.##
.#.
`

var c1z1 = `
#..
..#
.#.
`

// After 2 cycles:

var c2zn2 = `
.....
.....
..#..
.....
.....
`

var c2zn1 = `
..#..
.#..#
....#
.#...
.....
`

var c2z0 = `
##...
##...
#....
....#
.###.
`

var c2z1 = `
..#..
.#..#
....#
.#...
.....
`

var c2z2 = `
.....
.....
..#..
.....
.....
`

// After 3 cycles:

var c3zn2 = `
.......
.......
..##...
..###..
.......
.......
.......
`

var c3zn1 = `
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...
`

var c3z0 = `
...#...
.......
#......
.......
.....##
.##.#..
...#...
`

var c3z1 = `
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...
`

var c3z2 = `
.......
.......
..##...
..###..
.......
.......
.......
`

var c1 = map[[3]int]bool{}

var c2 = map[[3]int]bool{}

var c3 = map[[3]int]bool{}

func init() {

	// Build cycle 1 map
	for k, v := range parseInput(c1zn1, 0, 1, -1) {
		c1[k] = v
	}
	for k, v := range parseInput(c1z0, 0, 1, 0) {
		c1[k] = v
	}
	for k, v := range parseInput(c1z1, 0, 1, 1) {
		c1[k] = v
	}

	// Build cycle 2 map
	for k, v := range parseInput(c2zn2, -1, 0, -2) {
		c2[k] = v
	}
	for k, v := range parseInput(c2zn1, -1, 0, -1) {
		c2[k] = v
	}
	for k, v := range parseInput(c2z0, -1, 0, 0) {
		c2[k] = v
	}
	for k, v := range parseInput(c2z1, -1, 0, 1) {
		c2[k] = v
	}
	for k, v := range parseInput(c2z2, -1, 0, 2) {
		c2[k] = v
	}

	// Build cycle 3 map
	for k, v := range parseInput(c3zn2, -2, -1, -2) {
		c3[k] = v
	}
	for k, v := range parseInput(c3zn1, -2, -1, -1) {
		c3[k] = v
	}
	for k, v := range parseInput(c3z0, -2, -1, 0) {
		c3[k] = v
	}
	for k, v := range parseInput(c3z1, -2, -1, 1) {
		c3[k] = v
	}
	for k, v := range parseInput(c3z2, -2, -1, 2) {
		c3[k] = v
	}

}
