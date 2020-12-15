package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate([]int{13, 16, 0, 12, 15, 1.}, 30000000))
}

func calculate(input []int, iterations int) int {

	is := make([]int, iterations)
	copy(is, input)

	last := map[int]int{}
	secondLast := map[int]int{}

	for i := 0; i < len(input); i++ {
		if lastSeen, found := last[is[i]]; found {
			secondLast[is[i]] = lastSeen
		}
		last[is[i]] = i
	}

	for i := len(input); i < iterations; i++ {
		lastSeen, found := secondLast[is[i-1]]
		if found {
			is[i] = i - 1 - lastSeen
		} else {
			is[i] = 0
		}
		if lastSeen, found := last[is[i]]; found {
			secondLast[is[i]] = lastSeen
		}
		last[is[i]] = i
	}

	return is[iterations-1]

}
