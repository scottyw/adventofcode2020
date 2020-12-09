package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

func TestParseInput(t *testing.T) {
	data := parseInput(strings.Split(input, "\n"))
	assert.Equal(t, 20, len(data))
}

func TestFindSequence(t *testing.T) {
	data := parseInput(strings.Split(input, "\n"))
	assert.Equal(t, 62, findSequence(data, 127))
}

func TestAttack(t *testing.T) {
	data := parseInput(strings.Split(input, "\n"))
	assert.Equal(t, 127, attack(data, 5))
}

func TestValidate(t *testing.T) {
	assert.Equal(t, false, validate([]int{2, 4, 1, 3, 5}, 1))
	assert.Equal(t, false, validate([]int{2, 4, 1, 3, 5}, 2))
	assert.Equal(t, true, validate([]int{2, 4, 1, 3, 5}, 3))
	assert.Equal(t, false, validate([]int{2, 4, 1, 3, 5}, 10))
	assert.Equal(t, true, validate([]int{2, 4, 1, 3, 5}, 9))
	assert.Equal(t, false, validate([]int{95, 102, 117, 150, 182}, 127))
	assert.Equal(t, true, validate([]int{65, 95, 102, 117, 150}, 182))
}
