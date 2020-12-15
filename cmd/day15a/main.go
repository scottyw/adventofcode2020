package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate([]int{13, 16, 0, 12, 15, 1.}))
}

func calculate(input []int) int {

	is := make([]int, 2020)
	copy(is, input)

	for i := len(input); i < 2020; i++ {
		var value int
	inner:
		for j := i - 2; j >= 0; j-- {
			if is[j] == is[i-1] {
				value = i - 1 - j
				break inner
			}
		}
		is[i] = value
	}

	return is[2019]

}
