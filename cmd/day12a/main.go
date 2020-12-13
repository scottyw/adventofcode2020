package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileLinesToStringSlice("input/12.txt")
	fmt.Println(findDistance(input))
}

func findDistance(steps []string) int {
	var x, y int
	rotation := 90 // Start facing east
	for _, step := range steps {

		// Parse the step
		if step == "" {
			continue
		}
		op := step[0]
		arg := aoc.Atoi(step[1:])

		// Move the ship
		switch op {
		case 'L':
			rotation = (rotation + 360 - (arg % 360)) % 360
		case 'R':
			rotation = (rotation + arg) % 360
		case 'F':
			switch rotation {
			case 0:
				y += arg
			case 180:
				y -= arg
			case 270:
				x -= arg
			case 90:
				x += arg
			default:
				panic(rotation)
			}
		case 'N':
			y += arg
		case 'S':
			y -= arg
		case 'W':
			x -= arg
		case 'E':
			x += arg
		default:
			panic(op)
		}

	}

	// Calulcate Manhattan distance
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y

}
