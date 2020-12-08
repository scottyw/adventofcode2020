package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

type op struct {
	ins string
	arg int
}

func main() {
	input := aoc.FileLinesToStringSlice("input/08.txt")
	ops := parseProgram(input)
	acc := detectLoop(ops)
	fmt.Println(acc)
}

func detectLoop(ops []op) int {
	acc := 0
	pc := 0
	loopDetector := make([]bool, len(ops))
	for {

		// Has the program counter been set to this value before?
		if loopDetector[pc] {
			return acc
		}

		// Update the loop detector with the current program counter
		loopDetector[pc] = true

		// Fetch and execute the operation at the current program counter
		o := ops[pc]
		switch o.ins {
		case "nop":
			pc++
		case "acc":
			{
				acc += o.arg
				pc++
			}
		case "jmp":
			pc += o.arg
		}

	}
}

func parseProgram(lines []string) []op {
	ops := []op{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			panic(fmt.Sprintf("malformed line: %s", line))
		}
		ops = append(ops, op{ins: tokens[0], arg: aoc.Atoi(tokens[1])})
	}
	return ops
}
