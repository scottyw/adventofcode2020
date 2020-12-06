package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileToString("input/06.txt")
	fmt.Println(totalCount(parseInput(input)))
}

func parseInput(input string) [][]string {
	records := strings.Split(input, "\n\n")
	groups := [][]string{}
	for _, record := range records {
		members := strings.Split(record, "\n")
		groups = append(groups, members)
	}
	return groups
}

func count(group []string) int {
	answers := map[rune]bool{}
	for _, answer := range group {
		for _, r := range answer {
			if r >= 'a' && r <= 'z' {
				answers[r] = true
			}
		}
	}
	return len(answers)
}

func totalCount(groups [][]string) int {
	var total int
	for _, group := range groups {
		total += count(group)
	}
	return total
}
