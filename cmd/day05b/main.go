package main

import (
	"fmt"
	"sort"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	seats := aoc.FileLinesToStringSlice("input/05.txt")
	sort.Strings(seats)
	fmt.Println(findSeat(seats))
}

func findSeat(seats []string) int {
	ids := fmap(seats, findSeatID)
	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i] == ids[i-1]+2 {
			return ids[i] - 1
		}
	}
	return 0
}

func findSeatID(s string) int {
	// The ID is really a binary number where F and L represent 0 and B and R represent 1
	var id int
	// For each letter in the ID ...
	for _, c := range s {
		// ... shift left to make space for the new binary digit ...
		id <<= 1
		// ... and set it to 1 if necessary
		if c == 'B' || c == 'R' {
			id |= 1
		}
	}
	return id
}

func fmap(strs []string, f func(string) int) []int {
	is := make([]int, len(strs))
	for i, v := range strs {
		is[i] = f(v)
	}
	return is
}
