package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConstraints = map[string]constraint{
	"class": {1, 3, 5, 7},
	"row":   {6, 11, 33, 44},
	"seat":  {13, 40, 45, 50},
}

var myTestTicket = []int{7, 1, 14}

var nearbyTestTickets = [][]int{
	{7, 3, 47},
	{40, 4, 50},
	{55, 2, 20},
	{38, 6, 12},
}

func TestCheckConstraints(t *testing.T) {
	assert.Equal(t, true, checkConstraints(testConstraints, 40))
	assert.Equal(t, true, checkConstraints(testConstraints, 2))
	assert.Equal(t, false, checkConstraints(testConstraints, 4))
	assert.Equal(t, false, checkConstraints(testConstraints, 12))
}

func TestCheckNearbyTicket(t *testing.T) {
	assert.Equal(t, 0, checkNearbyTicket(testConstraints, nearbyTestTickets[0]))
	assert.Equal(t, 4, checkNearbyTicket(testConstraints, nearbyTestTickets[1]))
	assert.Equal(t, 55, checkNearbyTicket(testConstraints, nearbyTestTickets[2]))
	assert.Equal(t, 12, checkNearbyTicket(testConstraints, nearbyTestTickets[3]))
}

func TestCheckNearbyTickets(t *testing.T) {
	assert.Equal(t, 71, checkNearbyTickets(testConstraints, nearbyTestTickets))
}
