package main

import (
	"fmt"
	"math/bits"
	"regexp"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileToString("input/20.txt")
	tiles := parseTiles(input)
	fmt.Println(findCorners(tiles))
}

func findCorners(tiles map[int][4]uint) int {

	// Build a map that lists all tiles with a particular edge
	matches := map[uint][]int{}
	for id, edges := range tiles {
		for _, edge := range edges {
			ne := normalize(edge)
			matches[ne] = append(matches[ne], id)
		}
	}

	// Count how many unique edges are associated with each ID
	counts := map[int]int{}
	for _, ids := range matches {
		if len(ids) == 1 {
			counts[ids[0]]++
		}
	}

	// We're looking for tiles with 2 edges unique to them and
	// we cross-check we found exactly four of them
	var corners []int
	sum := 1
	for id, count := range counts {
		if count == 2 {
			sum *= id
			corners = append(corners, id)
		}
	}
	if len(corners) != 4 {
		panic(corners)
	}
	return sum

}

// Normalize an 10-bit edge which can be read forwards or backwards
func normalize(e uint) uint {
	r := uint(bits.Reverse16(uint16(e)) >> 6)
	if r > e {
		return e
	}
	return r
}

func parseTiles(input string) map[int][4]uint {
	tiles := map[int][4]uint{}
	for _, tile := range strings.Split(strings.TrimSpace(input), "\n\n") {
		id, edges := parseTile(strings.Split(tile, "\n"))
		tiles[id] = edges
	}
	return tiles
}

func parseTile(tile []string) (int, [4]uint) {

	if len(tile) != 11 {
		panic(fmt.Sprintf("unexpected tile: %v", tile))
	}

	matches := r.FindStringSubmatch(tile[0])
	id := aoc.Atoi(matches[1])

	// Convert top edge to uint
	var a uint
	for _, c := range tile[1] {
		a <<= 1
		if c == '#' {
			a |= 1
		}
	}

	// Convert bottom edge to uint
	var b uint
	for _, c := range tile[10] {
		b <<= 1
		if c == '#' {
			b |= 1
		}
	}

	// Convert left edge to uint
	var l uint
	for _, c := range tile[1:] {
		l <<= 1
		if c[0] == '#' {
			l |= 1
		}
	}

	// Convert right edge to uint
	var r uint
	for _, c := range tile[1:] {
		r <<= 1
		if c[9] == '#' {
			r |= 1
		}
	}

	return id, [4]uint{a, b, l, r}

}

var r = regexp.MustCompile("Tile (\\d+):")
