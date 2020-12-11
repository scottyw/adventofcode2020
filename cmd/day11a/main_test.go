package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var step0 = `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

var step1 = `
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`
var step2 = `
#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
`

func TestParseInput(t *testing.T) {
	cells := parseInput(strings.Split(step0, "\n"))
	assert.Equal(t, 10, len(cells))
	assert.Equal(t, 10, len(cells[0]))
}

func TestStabilize(t *testing.T) {
	cells := parseInput(strings.Split(step0, "\n"))
	assert.Equal(t, 37, stabilize(cells))
}

func TestIterate(t *testing.T) {
	cells0 := parseInput(strings.Split(step0, "\n"))
	cells1 := parseInput(strings.Split(step1, "\n"))
	cells2 := parseInput(strings.Split(step2, "\n"))
	assert.Equal(t, cells1, iterate(cells0))
	assert.Equal(t, cells2, iterate(cells1))
}
