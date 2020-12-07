package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

func TestParseRule(t *testing.T) {
	assert.Equal(t, 2, len(parseRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")))
	assert.Equal(t, 1, len(parseRule("bright white bags contain 1 shiny gold bag.")))
	assert.Equal(t, 0, len(parseRule("faded blue bags contain no other bags.")))
}

func TestBuildRulesTree(t *testing.T) {
	rules := buildRulesTree(strings.Split(input, "\n"))
	assert.Equal(t, 7, len(rules))
	assert.Equal(t, 2, len(rules["bright white"]))
	assert.Equal(t, 3, len(rules["faded blue"]))
}

func TestTraverseRulesTree(t *testing.T) {
	rules := buildRulesTree(strings.Split(input, "\n"))
	assert.Equal(t, 4, len(traverseRulesTree(rules, "shiny gold")))
}
