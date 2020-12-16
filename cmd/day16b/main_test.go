package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConstraints = []constraint{
	{"class", 0, 1, 4, 19},
	{"row", 0, 5, 8, 19},
	{"seat", 0, 13, 16, 19},
}

var myTestTicket = []int{11, 12, 13}

var nearbyTestTickets = [][]int{
	{3, 9, 18},
	{15, 1, 5},
	{5, 14, 9},
}

func TestDetermineFields(t *testing.T) {
	assert.Equal(t, []string{"row", "class", "seat"}, determineFields(testConstraints, nearbyTestTickets))
}

func TestValidateTicket(t *testing.T) {
	assert.Nil(t, validateTicket(testConstraints, []int{3, 9, 20}))
	assert.Equal(t, []matchedConstraints{6, 7, 7}, validateTicket(testConstraints, nearbyTestTickets[0]))
	assert.Equal(t, []matchedConstraints{3, 7, 7}, validateTicket(testConstraints, nearbyTestTickets[1]))
	assert.Equal(t, []matchedConstraints{7, 3, 7}, validateTicket(testConstraints, nearbyTestTickets[2]))
}

func TestMatchConstraints(t *testing.T) {
	assert.Equal(t, matchedConstraints(7), matchConstraints(testConstraints, 0))
	assert.Equal(t, matchedConstraints(6), matchConstraints(testConstraints, 3))
	assert.Equal(t, matchedConstraints(5), matchConstraints(testConstraints, 7))
	assert.Equal(t, matchedConstraints(7), matchConstraints(testConstraints, 10))
	assert.Equal(t, matchedConstraints(7), matchConstraints(testConstraints, 19))
	assert.Equal(t, matchedConstraints(0), matchConstraints(testConstraints, 20))
}
