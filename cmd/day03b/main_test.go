package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestCountTrees(t *testing.T) {

	trees := strings.Split(input, "\n")
	assert.Equal(t, 2, countTrees(trees, 1, 1))
	assert.Equal(t, 7, countTrees(trees, 3, 1))
	assert.Equal(t, 3, countTrees(trees, 5, 1))
	assert.Equal(t, 4, countTrees(trees, 7, 1))
	assert.Equal(t, 2, countTrees(trees, 1, 2))

}
