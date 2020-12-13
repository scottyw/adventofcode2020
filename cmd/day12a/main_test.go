package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
F10
N3
F7
R90
F11
`

func TestFindDistance(t *testing.T) {
	assert.Equal(t, 25, findDistance(strings.Split(input, "\n")))
}
