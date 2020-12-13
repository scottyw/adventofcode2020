package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileToString("input/13.txt")
	buses := parseInput(input)
	fmt.Println(findTime(buses))
}

func parseInput(input string) []int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	if len(lines) != 2 {
		panic("to many lines")
	}
	var buses []int
	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			buses = append(buses, 0)
		} else {
			buses = append(buses, aoc.Atoi(bus))
		}
	}
	return buses
}

func findTime(buses []int) int {

	// Start with the first bus which handily is index 0 in our input
	len := len(buses)
	x := buses[0]
	dx := len

	// We actually compute the time the last bus departs and treat every other bus as "integer N * ID + offset"
	for i := 1; i < len; i++ {
		if buses[i] == 0 {
			continue
		}
		// Find the LCM and offset for each pair of buses and use that as input in the next iteration with the next available bus
		x, dx = iterate(x, dx, buses[i], len-i)
	}

	return dx - len

}

// This function finds the LCM and the lowest value for which Mx + dx == Ny + dy when M and N are integers
func iterate(x, dx, y, dy int) (int, int) {

	// Find the lowest common multiple relying on the fact that all inputs are primes!
	lcm := x * y

	// Find the lowest value where Mx + dx == Ny + dy
	for i := 1; i < lcm; i++ {
		xs := (i * x) + dx
		if ((xs - dy) % y) == 0 {
			return lcm, xs
		}
	}

	// We should always find a match
	panic("iterate")

}
