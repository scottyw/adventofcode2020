package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `
16
10
15
5
1
11
7
19
6
12
4
`

var input2 = `
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

func TestFindChain(t *testing.T) {
	data := parseInput(strings.Split(input, "\n"))
	assert.Equal(t, 8, countCombinations(data))
	data2 := parseInput(strings.Split(input2, "\n"))
	assert.Equal(t, 19208, countCombinations(data2))
}

func TestParseInput(t *testing.T) {
	data := parseInput(strings.Split(input, "\n"))
	assert.Equal(t, 11, len(data))
}
