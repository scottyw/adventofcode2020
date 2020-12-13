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
	dx := 10 // Waypoint is 10 units east
	dy := 1  // Waypoint is 1  units north

	for _, step := range steps {

		// Parse the step
		if step == "" {
			continue
		}
		op := step[0]
		arg := aoc.Atoi(step[1:])

		// Move the ship
		switch op {
		case 'L', 'R':
			dx, dy = updateWaypoint(dx, dy, op, arg)
		case 'F':
			for i := 0; i < arg; i++ {
				x += dx
				y += dy
			}
		case 'N':
			dy += arg
		case 'S':
			dy -= arg
		case 'W':
			dx -= arg
		case 'E':
			dx += arg
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

func updateWaypoint(dx, dy int, op byte, arg int) (int, int) {

	// Normalize arg
	arg = arg % 360

	// Convert anti-clockwise into clockwise rotation
	if op == 'L' {
		arg = 360 - arg
	}

	// Rotate the waypoint clockwise
	switch arg {
	case 0:
		return dx, dy
	case 90:
		return dy, -dx
	case 180:
		return -dx, -dy
	case 270:
		return -dy, dx
	default:
		panic(arg)
	}

}
