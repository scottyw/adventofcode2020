package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileLinesToStringSlice("input/09.txt")
	data := parseInput(input)
	target := attack(data, 25)
	fmt.Println(findSequence(data, target))
}

func findSequence(data []int, candidate int) int {
	i := 0
	j := 1
	total := data[i] + data[j]
	for j < len(data) {
		switch {
		case total < candidate:
			j++
			total += data[j]
		case total > candidate:
			total -= data[i]
			i++
		case total == candidate:
			min := data[i]
			max := data[i]
			for _, x := range data[i : j+1] {
				if x > max {
					max = x
				}
				if x < min {
					min = x
				}
			}
			return min + max
		}
	}
	return 0
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
