package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1-3 a: abcde - valid
// 1-3 b: cdefg - invalid
// 2-9 c: ccccccccc - valid

func TestIsValidPassword(t *testing.T) {

	assert.True(t, isValidPassword("1-3 a: abcde"))
	assert.False(t, isValidPassword("1-3 b: cdefg"))
	assert.True(t, isValidPassword("2-9 c: ccccccccc"))

}
