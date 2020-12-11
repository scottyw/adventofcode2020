package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

const (
	empty uint8 = iota
	floor
	occupied
)

func main() {
	input := aoc.FileLinesToStringSlice("input/11.txt")
	cells := parseInput(input)
	fmt.Println(stabilize(cells))
}

func stabilize(current [][]uint8) int {
	var previous [][]uint8
	for !assert.ObjectsAreEqual(current, previous) {
		previous = current
		current = iterate(current)
	}
	return countOccupied(current)
}

func countOccupied(current [][]uint8) int {
	var occupiedCount int
	for _, row := range current {
		for _, cell := range row {
			if cell == occupied {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}

func iterate(current [][]uint8) [][]uint8 {
	next := [][]uint8{}
	mx := len(current)
	my := len(current[0])
	for x, row := range current {
		nextRow := []uint8{}
		for y, cell := range row {
			var occupiedCount int

			// Check 3 neighbours on the left
			if x > 0 && y > 0 && current[x-1][y-1] == occupied {
				occupiedCount++
			}
			if x > 0 && current[x-1][y] == occupied {
				occupiedCount++
			}
			if x > 0 && y < my-1 && current[x-1][y+1] == occupied {
				occupiedCount++
			}

			// Check 3 neighbours on the right
			if x < mx-1 && y > 0 && current[x+1][y-1] == occupied {
				occupiedCount++
			}
			if x < mx-1 && current[x+1][y] == occupied {
				occupiedCount++
			}
			if x < mx-1 && y < my-1 && current[x+1][y+1] == occupied {
				occupiedCount++
			}

			// Check 2 neighbours above and below
			if y > 0 && current[x][y-1] == occupied {
				occupiedCount++
			}
			if y < my-1 && current[x][y+1] == occupied {
				occupiedCount++
			}

			if cell == empty && occupiedCount == 0 {
				// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
				nextRow = append(nextRow, occupied)
			} else if cell == occupied && occupiedCount >= 4 {
				// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
				nextRow = append(nextRow, empty)
			} else {
				// Otherwise, the seat's state does not change.
				nextRow = append(nextRow, cell)
			}
		}
		next = append(next, nextRow)
	}
	return next
}

func parseInput(lines []string) [][]uint8 {
	cells := [][]uint8{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		row := []uint8{}
		for _, r := range line {
			switch r {
			case '.':
				row = append(row, floor)
			case 'L':
				row = append(row, empty)
			case '#':
				row = append(row, occupied)
			default:
				panic("parsing input")
			}
		}
		cells = append(cells, row)
	}
	return cells
}
