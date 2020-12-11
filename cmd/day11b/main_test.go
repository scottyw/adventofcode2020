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
#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#
`

var step3 = `
#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#
`

var occupied8 = `
.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....
`

var occupied0 = `
.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.
`

func TestParseInput(t *testing.T) {
	cells := parseInput(strings.Split(step0, "\n"))
	assert.Equal(t, 10, len(cells))
	assert.Equal(t, 10, len(cells[0]))
}

func TestCountCellOccupied(t *testing.T) {
	assert.Equal(t, 0, countCellOccupied(parseInput(strings.Split(occupied0, "\n")), 3, 3))
	assert.Equal(t, 8, countCellOccupied(parseInput(strings.Split(occupied8, "\n")), 3, 4))
}

func TestStabilize(t *testing.T) {
	cells := parseInput(strings.Split(step0, "\n"))
	assert.Equal(t, 26, stabilize(cells))
}

func TestIterate(t *testing.T) {
	cells0 := parseInput(strings.Split(step0, "\n"))
	cells1 := parseInput(strings.Split(step1, "\n"))
	cells2 := parseInput(strings.Split(step2, "\n"))
	cells3 := parseInput(strings.Split(step3, "\n"))
	assert.Equal(t, cells1, iterate(cells0), render(iterate(cells0)))
	assert.Equal(t, cells2, iterate(cells1), render(iterate(cells1)))
	assert.Equal(t, cells3, iterate(cells2), render(iterate(cells2)))
}
