package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileToString("input/13.txt")
	time, buses := parseInput(input)
	fmt.Println(findBus(time, buses))
}

func parseInput(input string) (int, []int) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	if len(lines) != 2 {
		panic("to many lines")
	}
	time := aoc.Atoi(lines[0])
	var buses []int
	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}
		buses = append(buses, aoc.Atoi(bus))
	}
	return time, buses
}

func findBus(time int, buses []int) int {
	nextBus := 0
	nextWait := math.MaxInt32
	for _, bus := range buses {
		wait := bus - (time % bus)
		if wait < nextWait {
			nextBus = bus
			nextWait = wait
		}
	}
	return nextBus * nextWait
}
