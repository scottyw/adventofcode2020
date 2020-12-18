package main

import (
	"fmt"

	"github.com/alecthomas/participle"
	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

var parser = participle.MustBuild(&Expression{})

func main() {
	lines := aoc.FileLinesToStringSlice("input/18.txt")
	var total int
	for _, line := range lines {
		total += eval(line)
	}
	fmt.Println(total)
}

func eval(s string) int {
	expr := &Expression{}
	err := parser.ParseString(s, expr)
	if err != nil {
		panic(err)
	}
	return expr.eval()
}

func (e *Expression) eval() int {
	result := e.Left.eval()
	for _, r := range e.Right {
		result = r.eval(result)
	}
	return result
}

func (ot *OpTerm) eval(l int) int {
	switch ot.Operator {
	case "+":
		return l + ot.Term.eval()
	case "*":
		return l * ot.Term.eval()
	default:
		panic("operator")
	}
}

func (t *Term) eval() int {
	if t.Number != nil {
		return *t.Number
	}
	return t.Subexpression.eval()
}

// Expression captures the whole input
type Expression struct {
	Left  *Term     `@@`
	Right []*OpTerm `@@*`
}

// OpTerm captures an Operator and a Term like "+ 2"
type OpTerm struct {
	Operator string `@("+" | "*")`
	Term     *Term  `@@`
}

// Term capture
type Term struct {
	Number        *int        `  @Int`
	Subexpression *Expression `| "(" @@ ")"`
}
