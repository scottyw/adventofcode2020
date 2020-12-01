package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	xs := aoc.FileLinesToIntSlice("input/01.txt")
	fmt.Println(findMatch(xs))
}

// We are brute forcing this because it is day 1 and we can
func findMatch(expenses []int) int {
	for i, x := range expenses {
		for j, y := range expenses {
			for k, z := range expenses {
				if i >= j || j >= k {
					continue
				}
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return 0
}
