package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
939
7,13,x,x,59,x,31,19
`

func TestParseInput(t *testing.T) {
	time, buses := parseInput(input)
	assert.Equal(t, 939, time)
	assert.Equal(t, buses, []int{7, 13, 59, 31, 19})
}

func TestFindBus(t *testing.T) {
	time, buses := parseInput(input)
	assert.Equal(t, 295, findBus(time, buses))
}
