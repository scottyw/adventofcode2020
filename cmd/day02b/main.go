package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

// 1-3 a: abcde
var r = regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")

func main() {
	lines := aoc.FileLinesToStringSlice("input/02.txt")

	count := 0
	for _, line := range lines {
		if isValidPassword(line) {
			count++
		}
	}
	fmt.Println(count)

}

func isValidPassword(line string) bool {
	matches := r.FindStringSubmatch(line)
	x := aoc.Atoi(matches[1])
	y := aoc.Atoi(matches[2])
	letter := matches[3]
	password := matches[4]
	return (password[x-1] == letter[0]) != (password[y-1] == letter[0])
}
