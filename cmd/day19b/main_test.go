package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Here, because rule 4 matches a and rule 5 matches b, rule 2 matches two letters that are the same (aa or bb), and rule 3 matches two letters that are different (ab or ba).

// Since rule 1 matches rules 2 and 3 once each in either order, it must match two pairs of letters, one pair with matching letters and one pair with different letters. This leaves eight possibilities: aaab, aaba, bbab, bbba, abaa, abbb, baaa, or babb.

// Rule 0, therefore, matches a (rule 4), then any of the eight options from rule 1, then b (rule 5): aaaabb, aaabab, abbabb, abbbab, aabaab, aabbbb, abaaab, or ababbb.

func TestMatch(t *testing.T) {
	assert.ElementsMatch(t, testResult, traverse(testRules, 0))
}

var testResult = []string{"aaaabb", "aaabab", "abbabb", "abbbab", "aabaab", "aabbbb", "abaaab", "ababbb"}

var testRules = [][][]int{
	{{4, 1, 5}},
	{{2, 3}, {3, 2}},
	{{4, 4}, {5, 5}},
	{{4, 5}, {5, 4}},
	{{-1}},
	{{-2}},
}
