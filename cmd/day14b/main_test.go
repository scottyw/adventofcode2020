package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`

func TestFindMaskedAddresses(t *testing.T) {
	maskedAddrs := findMaskedAddresses(42, "000000000000000000000000000000X1001X")
	assert.ElementsMatch(t, []uint64{26, 27, 58, 59}, maskedAddrs)
}

func TestReadInstruction(t *testing.T) {
	addr, value := readInstruction("mem[7] = 101")
	assert.Equal(t, uint64(7), addr)
	assert.Equal(t, uint64(101), value)
}

func TestRun(t *testing.T) {
	assert.Equal(t, uint64(208), sum(run(strings.Split(input, "\n"))))
}
