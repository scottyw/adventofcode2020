package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSeatID(t *testing.T) {

	// BFFFBBFRRR: row 70, column 7, seat ID 567.
	// FFFBBBFRRR: row 14, column 7, seat ID 119.
	// BBFFBBFRLL: row 102, column 4, seat ID 820.

	assert.Equal(t, 567, findSeatID("BFFFBBFRRR"))
	assert.Equal(t, 119, findSeatID("FFFBBBFRRR"))
	assert.Equal(t, 820, findSeatID("BBFFBBFRLL"))

}

func TestMaxSeatID(t *testing.T) {
	assert.Equal(t, 820, maxSeatID([]string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}))
}
