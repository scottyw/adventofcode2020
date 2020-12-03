package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	trees := aoc.FileLinesToStringSlice("input/03.txt")
	fmt.Println(countTrees(trees))
}

func countTrees(trees []string) int {
	var x, count int
	width := len(trees[0])
	for y := 0; y < len(trees); y++ {
		if trees[y][x] == '#' {
			count++
		}
		x = (x + 3) % width
	}
	return count
}
