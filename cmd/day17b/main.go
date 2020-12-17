package main

import (
	"fmt"
	"strings"
)

var input = `
####.#..
.......#
#..#####
.....##.
##...###
#..#.#.#
.##...#.
#...##..
`

func main() {
	cells := parseInput(input)
	for i := 0; i < 6; i++ {
		cells = iterate(cells)
	}
	fmt.Println(len(cells))
}

func iterate(current map[[4]int]bool) map[[4]int]bool {
	activeNeighbourCounts := map[[4]int]int{}
	for k := range current {

		x := k[0]
		y := k[1]
		z := k[2]
		w := k[3]

		for ix := x - 1; ix <= x+1; ix++ {
			for iy := y - 1; iy <= y+1; iy++ {
				for iz := z - 1; iz <= z+1; iz++ {
					for iw := w - 1; iw <= w+1; iw++ {
						activeNeighbourCounts[[4]int{ix, iy, iz, iw}]++
					}
				}
			}
		}

		// We've incremented x,y,z,w in the above for loop, so undo that now
		activeNeighbourCounts[[4]int{x, y, z, w}]--

	}

	next := map[[4]int]bool{}
	for k, activeNeighbourCount := range activeNeighbourCounts {
		_, active := current[k]
		if active && (activeNeighbourCount == 2 || activeNeighbourCount == 3) {
			next[k] = true
		}
		if !active && activeNeighbourCount == 3 {
			next[k] = true
		}
	}
	return next

}

func parseInput(input string) map[[4]int]bool {
	var y int
	cells := map[[4]int]bool{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for x, r := range line {
			switch r {
			case '.':
				// Ignore
			case '#':
				cells[[4]int{x, y, 0, 0}] = true
			default:
				panic("parsing input")
			}
		}
		y++
	}
	return cells
}
