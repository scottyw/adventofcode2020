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
	assert.Equal(t, 7, countTrees(trees))

}
