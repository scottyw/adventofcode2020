package main

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

// Represents a set of matched constraints e.g. if contraints 0, 2 and 5 match then bits 0, 2 and 5 will be set
type matchedConstraints uint

type constraint struct {
	name  string
	low1  int
	high1 int
	low2  int
	high2 int
}

func main() {
	nearbyTickets := aoc.FileLinesToIntSliceSlice("input/16.txt")
	fields := determineFields(constraints, nearbyTickets)
	fmt.Println(multiplyDepartureFields(fields, myTicket))
}

func multiplyDepartureFields(fields []string, tickets []int) int {
	total := 1
	for i, f := range fields {
		if strings.HasPrefix(f, "departure") {
			total *= tickets[i]
		}
	}
	return total
}

func determineFields(constraints []constraint, tickets [][]int) []string {

	// Find all valid tickets
	var validTickets [][]matchedConstraints
	for _, ticket := range tickets {
		matchedConstraints := validateTicket(constraints, ticket)
		if matchedConstraints != nil {
			validTickets = append(validTickets, matchedConstraints)
		}
	}

	// For each field, look at all the valid tickets to see which constraints match universally
	fieldCount := len(constraints)
	universalMatches := make([]matchedConstraints, fieldCount)
	for i := 0; i < fieldCount; i++ {
		universalMatch := matchedConstraints((1 << (fieldCount + 1)) - 1)
		for _, validTicket := range validTickets {
			universalMatch &= validTicket[i]
		}
		universalMatches[i] = universalMatch
	}

	// Work out which field matches which index in the ticket
	fields := make([]string, fieldCount)
	for {

		// Look for a field that matched exactly one constraint universally i.e. has just one bit set
		matchIndex, err := findExactUniversalMatch(universalMatches)
		if err != nil {
			break
		}
		matchValue := universalMatches[matchIndex]

		// Unset the matched bit in all elements (including itself)
		for i := range universalMatches {
			universalMatches[i] &^= matchValue
		}

		// Find the actual constraint for the match
		var i int
		for matchValue > 1 {
			i++
			matchValue >>= 1
		}
		fields[matchIndex] = constraints[i].name

	}
	return fields

}

// Find an index into universal matches where the element has exactly one bit set
func findExactUniversalMatch(universalMatches []matchedConstraints) (int, error) {
	for i := range universalMatches {
		if bits.OnesCount(uint(universalMatches[i])) == 1 {
			return i, nil
		}
	}
	return 0, fmt.Errorf("no universal match")
}

func validateTicket(constraints []constraint, ticket []int) []matchedConstraints {
	matching := make([]matchedConstraints, len(ticket))
	for i, x := range ticket {
		match := matchConstraints(constraints, x)
		if match == 0 {
			// This field is invalid and so we discard this ticket
			return nil
		}
		matching[i] = match
	}
	return matching
}

// Check which constraints x matches
func matchConstraints(constraints []constraint, x int) matchedConstraints {
	var matchedConstraints matchedConstraints
	for i, constraint := range constraints {
		if (constraint.low1 <= x && constraint.high1 >= x) ||
			(constraint.low2 <= x && constraint.high2 >= x) {
			matchedConstraints |= 1 << i
		}
	}
	return matchedConstraints
}

var myTicket = []int{71, 127, 181, 179, 113, 109, 79, 151, 97, 107, 53, 193, 73, 83, 191, 101, 89, 149, 103, 197}

var constraints = []constraint{
	{"departure location", 41, 598, 605, 974},
	{"departure station", 30, 617, 625, 957},
	{"departure platform", 29, 914, 931, 960},
	{"departure track", 39, 734, 756, 972},
	{"departure date", 37, 894, 915, 956},
	{"departure time", 48, 54, 70, 955},
	{"arrival location", 39, 469, 491, 955},
	{"arrival station", 47, 269, 282, 949},
	{"arrival platform", 26, 500, 521, 960},
	{"arrival track", 26, 681, 703, 953},
	{"class", 49, 293, 318, 956},
	{"duration", 25, 861, 873, 973},
	{"price", 30, 446, 465, 958},
	{"route", 50, 525, 551, 973},
	{"row", 39, 129, 141, 972},
	{"seat", 37, 566, 573, 953},
	{"train", 43, 330, 356, 969},
	{"type", 32, 770, 792, 955},
	{"wagon", 47, 435, 446, 961},
	{"zone", 30, 155, 179, 957},
}
