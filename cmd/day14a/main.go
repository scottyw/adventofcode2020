package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

type instruction struct {
	addr  uint64
	value uint64
}

var instructionRegex = regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")

func main() {
	program := aoc.FileLinesToStringSlice("input/14.txt")
	mem := run(program)
	fmt.Println(sum(mem))
}

func sum(mem map[uint64]uint64) uint64 {
	var total uint64
	for _, v := range mem {
		total += v
	}
	return total
}

func run(program []string) map[uint64]uint64 {
	var setMask, unsetMask uint64
	mem := map[uint64]uint64{}
	for _, instruction := range program {
		switch {
		case instruction == "":
			continue
		case strings.HasPrefix(instruction, "mask = "):
			setMask, unsetMask = updateMask(instruction)
		case strings.HasPrefix(instruction, "mem"):
			addr, value := readInstruction(instruction)
			value |= setMask
			value &^= unsetMask
			mem[addr] = value
		default:
			panic(instruction)
		}
	}
	return mem
}

func updateMask(instruction string) (uint64, uint64) {
	mask := strings.TrimPrefix(instruction, "mask = ")
	setMask, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)
	if err != nil {
		panic("mask set bits parsing")
	}
	mask = strings.ReplaceAll(mask, "1", "X")
	mask = strings.ReplaceAll(mask, "0", "1")
	mask = strings.ReplaceAll(mask, "X", "0")
	unsetMask, err := strconv.ParseUint(mask, 2, 64)
	if err != nil {
		panic("mask unset bits parsing")
	}
	return setMask, unsetMask
}

func readInstruction(instruction string) (uint64, uint64) {
	matches := instructionRegex.FindStringSubmatch(instruction)
	if matches == nil {
		panic("program line")
	}
	addr, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		panic(err)
	}
	value, err := strconv.ParseUint(matches[2], 10, 64)
	if err != nil {
		panic(err)
	}
	return addr, value
}
