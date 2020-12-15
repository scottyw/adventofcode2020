package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {

	// Given the starting numbers 1,3,2, the 2020th number spoken is 1.
	assert.Equal(t, 1, calculate([]int{1, 3, 2}))

	// Given the starting numbers 2,1,3, the 2020th number spoken is 10.
	assert.Equal(t, 10, calculate([]int{2, 1, 3}))

	// Given the starting numbers 1,2,3, the 2020th number spoken is 27.
	assert.Equal(t, 27, calculate([]int{1, 2, 3}))

	// Given the starting numbers 2,3,1, the 2020th number spoken is 78.
	assert.Equal(t, 78, calculate([]int{2, 3, 1}))

	// Given the starting numbers 3,2,1, the 2020th number spoken is 438.
	assert.Equal(t, 438, calculate([]int{3, 2, 1}))

	// Given the starting numbers 3,1,2, the 2020th number spoken is 1836.
	assert.Equal(t, 1836, calculate([]int{3, 1, 2}))

}
