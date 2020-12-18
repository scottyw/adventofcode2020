package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	assert.Equal(t, 71, eval("1 + 2 * 3 + 4 * 5 + 6"))
	assert.Equal(t, 26, eval("2 * 3 + (4 * 5)"))
	assert.Equal(t, 437, eval("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	assert.Equal(t, 12240, eval("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	assert.Equal(t, 13632, eval("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}
