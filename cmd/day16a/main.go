package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

type constraint struct {
	low1  int
	high1 int
	low2  int
	high2 int
}

func main() {
	nearbyTickets := aoc.FileLinesToIntSliceSlice("input/16.txt")
	fmt.Println(checkNearbyTickets(constraints, nearbyTickets))
}

func checkNearbyTickets(constraints map[string]constraint, tickets [][]int) int {
	var total int
	for _, ticket := range tickets {
		total += checkNearbyTicket(constraints, ticket)
	}
	return total
}

func checkNearbyTicket(constraints map[string]constraint, ticket []int) int {
	var total int
	for _, x := range ticket {
		if !checkConstraints(constraints, x) {
			total += x
		}
	}
	return total
}

func checkConstraints(constraints map[string]constraint, x int) bool {
	for _, constraint := range constraints {
		if (constraint.low1 <= x && constraint.high1 >= x) ||
			(constraint.low2 <= x && constraint.high2 >= x) {
			return true
		}
	}
	return false
}

var myTicket = []int{71, 127, 181, 179, 113, 109, 79, 151, 97, 107, 53, 193, 73, 83, 191, 101, 89, 149, 103, 197}

var constraints = map[string]constraint{
	"departure location": {41, 598, 605, 974},
	"departure station":  {30, 617, 625, 957},
	"departure platform": {29, 914, 931, 960},
	"departure track":    {39, 734, 756, 972},
	"departure date":     {37, 894, 915, 956},
	"departure time":     {48, 54, 70, 955},
	"arrival location":   {39, 469, 491, 955},
	"arrival station":    {47, 269, 282, 949},
	"arrival platform":   {26, 500, 521, 960},
	"arrival track":      {26, 681, 703, 953},
	"class":              {49, 293, 318, 956},
	"duration":           {25, 861, 873, 973},
	"price":              {30, 446, 465, 958},
	"route":              {50, 525, 551, 973},
	"row":                {39, 129, 141, 972},
	"seat":               {37, 566, 573, 953},
	"train":              {43, 330, 356, 969},
	"type":               {32, 770, 792, 955},
	"wagon":              {47, 435, 446, 961},
	"zone":               {30, 155, 179, 957},
}
