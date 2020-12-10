package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileLinesToStringSlice("input/10.txt")
	chain := parseInput(input)
	fmt.Println(countCombinations(chain))
}

/****************************************************************

From part A we know that the input is a chain with gaps of 1 or 3 but not 0, 2 or more than 3

Any adapter with a gap of 3 before or after cannot be optional or the gap would be too large if it were removed.

These ones will appear in every valid combination.

Any adaptor with a gap of 1 both before and after it is potentially optional.

We search for sequences of consective gaps of size 1.

In practice we don't need to check gaps on both sides of adapter.

We look at the previous gap only which is enough to work out how many valid combinations there are of each sequence

 - [0 3 4 5 8] contains sequence [4 5] of length 2
 - [0 3 4 5 6 9] contains sequence [4 5 6] of length 3
 - [0 3 4 5 6 7 10] contains sequence [4 5 6 7] of length 4
 - etc.

With this input we luckily only need to worry about N = 4 at most:
 - length = 1, valid combinations = 1 e.g. [1 4 5 7]
 - length = 2, valid combinations = 2 e.g. [1 4 5 6 9] [1 4 6 9]
 - length = 3, valid combinations = 4 e.g. [1 4 5 6 7 10] [1 4 5 7 10] [1 4 6 7 10] [1 4 7 10]
 - length = 4, valid combinations = 7

Once we have a count of the valid combinations for each sequence, we multiple them together to find the total count of possible combinations across the whole chain.

****************************************************************/

func countCombinations(adapters []int) int {

	// Add the socket on the plane to make gap calculation simpler
	adapters = append(adapters, 0)

	// Sort the adapters by joltage
	sort.Ints(adapters)

	// Add the adapter on the device too
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	// Implement the algorithm above
	var gap, consecutive int
	total := 1
	for i := 1; i < len(adapters); i++ {
		// Each adapter connects to the previous one
		gap = adapters[i] - adapters[i-1]
		switch gap {
		case 1:
			consecutive++
		case 3:
			if consecutive > 0 {
				switch consecutive {
				case 1:
					// Ignore
				case 2:
					total *= 2
				case 3:
					total *= 4
				case 4:
					total *= 7
				default:
					panic(fmt.Sprintf("consecutive: %d", consecutive))
				}
				consecutive = 0
			}
		default:
			panic(fmt.Sprintf("gap size: %d", i))
		}
	}

	return total
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
