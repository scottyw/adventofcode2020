package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

type node struct {
	count  int
	colour string
}

func main() {
	input := aoc.FileLinesToStringSlice("input/07.txt")
	rulesTree := buildRulesTree(input)
	fmt.Println(countBagContents(rulesTree, "shiny gold"))
}

func countBagContents(rulesTree map[string][]node, colour string) int {
	var count int
	for _, child := range rulesTree[colour] {
		count += child.count * (countBagContents(rulesTree, child.colour) + 1)
	}
	return count
}

func buildRulesTree(ruleStrs []string) map[string][]node {
	rules := map[string][]node{}
	for _, ruleStr := range ruleStrs {
		if ruleStr == "" {
			continue
		}
		rule := parseRule(strings.TrimSpace(ruleStr))
		for k, v := range rule {
			if _, ok := rules[k]; ok {
				panic("Two rules affect a single colour")
			}
			rules[k] = v
		}
	}
	return rules
}

// bright white bags contain 1 shiny gold bag.
// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
// faded blue bags contain no other bags.
var leftSideRegex = regexp.MustCompile("^(\\w+ \\w+) bags contain (.+)\\.$")
var rightSideRegex = regexp.MustCompile("(\\d+) (\\w+ \\w+) bag")

func parseRule(ruleStr string) map[string][]node {
	left := leftSideRegex.FindStringSubmatch(ruleStr)
	if left == nil {
		panic(fmt.Sprintf("no left side rule match: %s", ruleStr))
	}
	if left[2] == "no other bags" {
		return nil
	}
	rights := rightSideRegex.FindAllStringSubmatch(left[2], -1)
	if rights == nil {
		panic(fmt.Sprintf("no right side rule match: %s", left[2]))
	}
	rule := map[string][]node{}
	for _, right := range rights {
		rule[left[1]] = append(rule[left[1]], node{
			count:  aoc.Atoi(right[1]),
			colour: right[2],
		})
	}
	return rule
}
