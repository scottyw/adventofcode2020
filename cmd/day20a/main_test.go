package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTiles(t *testing.T) {
	tiles := parseTiles(testTiles)
	assert.Equal(t, 9, len(tiles))
	assert.Equal(t, [4]uint{0xd2, 0xe7, 0x1f2, 0x59}, tiles[2311])
}

func TestFindCorners(t *testing.T) {
	tiles := parseTiles(testTiles)
	assert.Equal(t, 20899048083289, findCorners(tiles))
}

func TestNormalize(t *testing.T) {
	assert.Equal(t, uint(0b0101010101), normalize(0b0101010101))
	assert.Equal(t, uint(0b0101010101), normalize(0b1010101010))
	assert.Equal(t, uint(0b0100110001), normalize(0b0100110001))
	assert.Equal(t, uint(0b0100110001), normalize(0b1000110010))
}

var testTiles = `

Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...


`
