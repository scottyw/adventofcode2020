package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	assert.Equal(t, 231, eval("1 + 2 * 3 + 4 * 5 + 6"))
	assert.Equal(t, 51, eval("1 + (2 * 3) + (4 * (5 + 6))"))
	assert.Equal(t, 46, eval("2 * 3 + (4 * 5)"))
	assert.Equal(t, 1445, eval("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	assert.Equal(t, 669060, eval("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	assert.Equal(t, 23340, eval("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}
