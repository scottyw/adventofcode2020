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

b`

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

	// The first group contains one person who answered "yes" to 3 questions: a, b, and c.
	// The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
	// The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
	// The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
	// The last group contains one person who answered "yes" to only 1 question, b.

	groups := parseInput(input)
	assert.Equal(t, 3, count(groups[0]))
	assert.Equal(t, 3, count(groups[1]))
	assert.Equal(t, 3, count(groups[2]))
	assert.Equal(t, 1, count(groups[3]))
	assert.Equal(t, 1, count(groups[4]))

}

func TestTotalCount(t *testing.T) {
	groups := parseInput(input)
	assert.Equal(t, 11, totalCount(groups))
}
