package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	trees := aoc.FileLinesToStringSlice("input/03.txt")
	total := countTrees(trees, 1, 1)
	total *= countTrees(trees, 3, 1)
	total *= countTrees(trees, 5, 1)
	total *= countTrees(trees, 7, 1)
	total *= countTrees(trees, 1, 2)
	fmt.Println(total)
}

func countTrees(trees []string, dx, dy int) int {
	var x, count int
	width := len(trees[0])
	for y := 0; y < len(trees); y += dy {
		if trees[y][x] == '#' {
			count++
		}
		x = (x + dx) % width
	}
	return count
}
