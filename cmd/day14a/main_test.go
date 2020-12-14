package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

func TestParseInput(t *testing.T) {
	setMask, unsetMask := updateMask("mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	assert.Equal(t, uint64(0x40), setMask)
	assert.Equal(t, uint64(0x02), unsetMask)
}

func TestReadInstruction(t *testing.T) {
	addr, value := readInstruction("mem[7] = 101")
	assert.Equal(t, uint64(7), addr)
	assert.Equal(t, uint64(101), value)
}

func TestRun(t *testing.T) {
	assert.Equal(t, uint64(165), sum(run(strings.Split(input, "\n"))))
}
