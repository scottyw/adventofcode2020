package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	xs := aoc.FileLinesToIntSlice("input/01.txt")
	fmt.Println(findMatch(xs))
}

func findMatch(expenses []int) int {
	for i, x := range expenses {
		for j, y := range expenses {
			if i >= j {
				continue
			}
			if x+y == 2020 {
				return x * y
			}
		}
	}
	return 0
}
