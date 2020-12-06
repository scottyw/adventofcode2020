package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `abc

a
b
c

ab
ac

a
a
a
a

b
`

func TestParseInput(t *testing.T) {
	groups := parseInput(input)
	assert.Equal(t, 5, len(groups))
	assert.Equal(t, 1, len(groups[0]))
	assert.Equal(t, 3, len(groups[1]))
	assert.Equal(t, 2, len(groups[2]))
	assert.Equal(t, 4, len(groups[3]))
	assert.Equal(t, 1, len(groups[4]))
}

func TestCount(t *testing.T) {

	// In the first group, everyone (all 1 person) answered "yes" to 3 questions: a, b, and c.
	// In the second group, there is no question to which everyone answered "yes".
	// In the third group, everyone answered yes to only 1 question, a. Since some people did not answer "yes" to b or c, they don't count.
	// In the fourth group, everyone answered yes to only 1 question, a.
	// In the fifth group, everyone (all 1 person) answered "yes" to 1 question, b.

	groups := parseInput(input)
	assert.Equal(t, 3, count(groups[0]))
	assert.Equal(t, 0, count(groups[1]))
	assert.Equal(t, 1, count(groups[2]))
	assert.Equal(t, 1, count(groups[3]))
	assert.Equal(t, 1, count(groups[4]))

}

func TestTotalCount(t *testing.T) {
	groups := parseInput(input)
	assert.Equal(t, 6, totalCount(groups))
}
