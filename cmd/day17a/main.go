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
	cells := parseInput(input, 0, 0, 0)
	for i := 0; i < 6; i++ {
		cells = iterate(cells)
	}
	fmt.Println(len(cells))
}

func iterate(current map[[3]int]bool) map[[3]int]bool {
	activeNeighbourCounts := map[[3]int]int{}
	for k := range current {

		x := k[0]
		y := k[1]
		z := k[2]

		activeNeighbourCounts[[3]int{x - 1, y - 1, z - 1}]++
		activeNeighbourCounts[[3]int{x - 1, y - 1, z}]++
		activeNeighbourCounts[[3]int{x - 1, y - 1, z + 1}]++

		activeNeighbourCounts[[3]int{x - 1, y, z - 1}]++
		activeNeighbourCounts[[3]int{x - 1, y, z}]++
		activeNeighbourCounts[[3]int{x - 1, y, z + 1}]++

		activeNeighbourCounts[[3]int{x - 1, y + 1, z - 1}]++
		activeNeighbourCounts[[3]int{x - 1, y + 1, z}]++
		activeNeighbourCounts[[3]int{x - 1, y + 1, z + 1}]++

		activeNeighbourCounts[[3]int{x, y - 1, z - 1}]++
		activeNeighbourCounts[[3]int{x, y - 1, z}]++
		activeNeighbourCounts[[3]int{x, y - 1, z + 1}]++

		activeNeighbourCounts[[3]int{x, y, z - 1}]++
		activeNeighbourCounts[[3]int{x, y, z + 1}]++

		activeNeighbourCounts[[3]int{x, y + 1, z - 1}]++
		activeNeighbourCounts[[3]int{x, y + 1, z}]++
		activeNeighbourCounts[[3]int{x, y + 1, z + 1}]++

		activeNeighbourCounts[[3]int{x + 1, y - 1, z - 1}]++
		activeNeighbourCounts[[3]int{x + 1, y - 1, z}]++
		activeNeighbourCounts[[3]int{x + 1, y - 1, z + 1}]++

		activeNeighbourCounts[[3]int{x + 1, y, z - 1}]++
		activeNeighbourCounts[[3]int{x + 1, y, z}]++
		activeNeighbourCounts[[3]int{x + 1, y, z + 1}]++

		activeNeighbourCounts[[3]int{x + 1, y + 1, z - 1}]++
		activeNeighbourCounts[[3]int{x + 1, y + 1, z}]++
		activeNeighbourCounts[[3]int{x + 1, y + 1, z + 1}]++

	}

	next := map[[3]int]bool{}
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

func parseInput(input string, dx, dy, dz int) map[[3]int]bool {
	var y int
	cells := map[[3]int]bool{}
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
				cells[[3]int{x + dx, y + dy, dz}] = true
			default:
				panic("parsing input")
			}
		}
		y++
	}
	return cells
}
