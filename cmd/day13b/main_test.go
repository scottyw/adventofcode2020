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
	buses := parseInput(input)
	assert.Equal(t, []int{7, 13, 0, 0, 59, 0, 31, 19}, buses)
}

func TestFindTime(t *testing.T) {
	buses := parseInput(input)
	assert.Equal(t, 1068781, findTime(buses))
}
