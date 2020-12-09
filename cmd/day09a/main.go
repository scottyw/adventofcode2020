package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileLinesToStringSlice("input/09.txt")
	data := parseInput(input)
	fmt.Println(attack(data, 25))
}

func attack(data []int, windowSize int) int {
	for i := windowSize; i < len(data); i++ {
		if !validate(data[i-windowSize:i], data[i]) {
			return data[i]
		}
	}
	return 0
}

func validate(window []int, candidate int) bool {
	for i, x := range window {
		for j, y := range window {
			if i >= j {
				continue
			}
			if x+y == candidate {
				return true
			}
		}
	}
	return false
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
