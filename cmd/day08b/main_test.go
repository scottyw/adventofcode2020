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

func TestFix(t *testing.T) {
	ops := parseProgram(strings.Split(input, "\n"))
	acc, err := fix(ops)
	assert.Nil(t, err)
	assert.Equal(t, 8, acc)
}
