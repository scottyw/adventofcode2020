package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

func TestParseProgram(t *testing.T) {
	ops := parseProgram(strings.Split(input, "\n"))
	assert.Equal(t, 9, len(ops))
}

func TestDetectLoop(t *testing.T) {
	ops := parseProgram(strings.Split(input, "\n"))
	acc := detectLoop(ops)
	assert.Equal(t, 5, acc)
}
