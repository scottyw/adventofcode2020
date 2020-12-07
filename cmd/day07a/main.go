package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	input := aoc.FileLinesToStringSlice("input/07.txt")
	rulesTree := buildRulesTree(input)
	fmt.Println(len(traverseRulesTree(rulesTree, "shiny gold")))
}

func traverseRulesTree(rulesTree map[string][]string, colour string) map[string]bool {
	colours := map[string]bool{}
	for _, child := range rulesTree[colour] {
		colours[child] = true
		for colour := range traverseRulesTree(rulesTree, child) {
			colours[colour] = true
		}
	}
	return colours
}

func buildRulesTree(ruleStrs []string) map[string][]string {
	rules := map[string][]string{}
	for _, ruleStr := range ruleStrs {
		if ruleStr == "" {
			continue
		}
		rule := parseRule(strings.TrimSpace(ruleStr))
		for k, v := range rule {
			rules[k] = append(rules[k], v)
		}
	}
	return rules
}

// bright white bags contain 1 shiny gold bag.
// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
// faded blue bags contain no other bags.
var leftSideRegex = regexp.MustCompile("^(\\w+ \\w+) bags contain (.+)\\.$")
var rightSideRegex = regexp.MustCompile("(\\w+ \\w+) bag")

func parseRule(ruleStr string) map[string]string {
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
	rule := map[string]string{}
	for _, right := range rights {
		rule[right[1]] = left[1]
	}
	return rule
}
