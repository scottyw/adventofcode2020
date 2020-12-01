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

// Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

func TestFindMatch(t *testing.T) {

	expenses := []int{1721, 979, 366, 299, 675, 1456}

	assert.Equal(t, 241861950, findMatch(expenses))

}
