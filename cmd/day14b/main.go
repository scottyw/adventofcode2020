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
	var mask string
	mem := map[uint64]uint64{}
	for _, instruction := range program {
		switch {
		case instruction == "":
			continue
		case strings.HasPrefix(instruction, "mask = "):
			mask = strings.TrimPrefix(instruction, "mask = ")
		case strings.HasPrefix(instruction, "mem"):
			addr, value := readInstruction(instruction)
			maskedAddrs := findMaskedAddresses(addr, mask)
			for _, maskedAddr := range maskedAddrs {
				mem[maskedAddr] = value
			}
		default:
			panic(instruction)
		}
	}
	return mem
}

func findMaskedAddresses(addr uint64, mask string) []uint64 {

	maskedAddrs := []uint64{0}
	l := len(mask) - 1

	// We're working with 36-bit numbers
	for i := 0; i < 36; i++ {
		switch mask[l-i] {
		case '0':
			// Use the value directly from 'addr'
			for j := range maskedAddrs {
				maskedAddrs[j] |= (addr & (1 << i))
			}
		case '1':
			// Set this bit in every address
			for j := range maskedAddrs {
				maskedAddrs[j] |= (1 << i)
			}
		case 'X':
			updatedAddrs := []uint64{}
			for _, maskedAddr := range maskedAddrs {
				updatedAddrs = append(updatedAddrs, maskedAddr&^(1<<i))
				updatedAddrs = append(updatedAddrs, maskedAddr|(1<<i))
			}
			maskedAddrs = updatedAddrs
		default:
			panic(mask[i])
		}
	}

	return maskedAddrs

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
