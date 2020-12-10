package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileLinesToStringSlice("input/10.txt")
	adapters := parseInput(input)
	fmt.Println(findChain(adapters))
}

func findChain(adapters []int) int {

	// Add the socket on the plane to make gap calculation simpler
	adapters = append(adapters, 0)

	// Sort the adapters by joltage
	sort.Ints(adapters)

	var joltage1, joltage3, gap int
	for i := 1; i < len(adapters); i++ {
		// Every other adapter connects to the previous one
		gap = adapters[i] - adapters[i-1]
		switch gap {
		case 1:
			joltage1++
		case 3:
			joltage3++
		default:
			panic(fmt.Sprintf("The gap is not manageable: %d", i))
		}
	}

	// And one more joltage3 for the built-in one ...
	joltage3++

	return joltage1 * joltage3

}

func parseInput(lines []string) []int {
	is := []int{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		is = append(is, aoc.Atoi(line))
	}
	return is
}
