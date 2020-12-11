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
	for y, row := range current {
		nextRow := []uint8{}
		for x, cell := range row {
			occupiedCount := countCellOccupied(current, x, y)
			if cell == empty && occupiedCount == 0 {
				// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
				nextRow = append(nextRow, occupied)
			} else if cell == occupied && occupiedCount >= 5 {
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

func countCellOccupied(current [][]uint8, x, y int) int {
	my := len(current)
	mx := len(current[0])

	var occupiedCount int

	// North
	if y > 0 {
		for dy := y - 1; dy >= 0; dy-- {
			if current[dy][x] == occupied {
				occupiedCount++
				break
			}
			if current[dy][x] == empty {
				break
			}
		}
	}

	// South
	if y < my-1 {
		for dy := y + 1; dy < my; dy++ {
			if current[dy][x] == occupied {
				occupiedCount++
				break
			}
			if current[dy][x] == empty {
				break
			}
		}
	}

	// West
	if x > 0 {
		for dx := x - 1; dx >= 0; dx-- {
			if current[y][dx] == occupied {
				occupiedCount++
				break
			}
			if current[y][dx] == empty {
				break
			}
		}
	}

	// East
	if x < mx-1 {
		for dx := x + 1; dx < mx; dx++ {
			if current[y][dx] == occupied {
				occupiedCount++
				break
			}
			if current[y][dx] == empty {
				break
			}
		}
	}

	// North West
	if x > 0 && y > 0 {
		dx := x - 1
		dy := y - 1
		for dx >= 0 && dy >= 0 {
			if current[dy][dx] == occupied {
				occupiedCount++
				break
			}
			if current[dy][dx] == empty {
				break
			}
			dx--
			dy--
		}
	}

	// North East
	if x < mx-1 && y > 0 {
		dx := x + 1
		dy := y - 1
		for dx < mx && dy >= 0 {
			if current[dy][dx] == occupied {
				occupiedCount++
				break
			}
			if current[dy][dx] == empty {
				break
			}
			dx++
			dy--
		}
	}

	// South East
	if x < mx-1 && y < my-1 {
		dx := x + 1
		dy := y + 1
		for dx < mx && dy < my {
			if current[dy][dx] == occupied {
				occupiedCount++
				break
			}
			if current[dy][dx] == empty {
				break
			}
			dx++
			dy++
		}
	}

	// South West
	if x > 0 && y < my-1 {
		dx := x - 1
		dy := y + 1
		for dx >= 0 && dy < my {
			if current[dy][dx] == occupied {
				occupiedCount++
				break
			}
			if current[dy][dx] == empty {
				break
			}
			dx--
			dy++
		}
	}

	// fmt.Printf("%d %d - %d\n", x, y, occupiedCount)

	return occupiedCount
}

func render(current [][]uint8) string {
	var s string
	for _, row := range current {
		for _, cell := range row {
			switch cell {
			case 0:
				s += "L"
			case 1:
				s += "."
			case 2:
				s += "#"
			default:
				panic("render")
			}
		}
		s += "\n"
	}
	return s
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
