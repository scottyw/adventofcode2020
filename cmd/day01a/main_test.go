package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For example, suppose your expense report contained the following:

// 1721
// 979
// 366
// 299
// 675
// 1456

// In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

func TestFindMatch(t *testing.T) {

	expenses := []int{1721, 979, 366, 299, 675, 1456}

	assert.Equal(t, 514579, findMatch(expenses))

}
