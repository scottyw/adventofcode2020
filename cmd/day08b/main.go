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
	acc, err := fix(ops)
	if err != nil {
		panic(err)
	}
	fmt.Println(acc)
}

func fix(ops []op) (int, error) {

	// Flip a single operation in the program and see if executes correctly
	for i := range ops {
		acc, err := executeWithFlippedOperation(ops, i)
		if err == nil {
			// Successfuol execution
			return acc, nil
		}
	}

	return 0, fmt.Errorf("no suitable instruction flip found")

}

func executeWithFlippedOperation(ops []op, pcToFlip int) (int, error) {
	acc := 0
	pc := 0
	loopDetector := make([]bool, len(ops))
	for {

		// Check for successful termination
		if pc >= len(ops) {
			return acc, nil
		}

		// Has the program counter been set to this value before?
		if loopDetector[pc] {
			return 0, fmt.Errorf("infinite loop detected")
		}

		// Update the loop detector with the current program counter
		loopDetector[pc] = true

		// Fetch and flip the instruction if it is the chosen one this time
		o := ops[pc]
		ins := o.ins
		if pc == pcToFlip {
			if ins == "nop" {
				ins = "jmp"
			} else if ins == "jmp" {
				ins = "nop"
			}
		}

		// Execute the (potentially flipped) operation
		switch ins {
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
