package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

var matches = traverse(rules, 0)

func main() {
	lines := aoc.FileLinesToStringSlice("input/19.txt")
	var total int
	for _, line := range lines {
		for _, m := range matches {
			if line == m {
				total++
				break
			}
		}
	}
	fmt.Println(total)
}

func traverse(rules [][][]int, i int) []string {
	var results []string
	for _, r := range rules[i] {
		xs := []string{""}
		for _, ri := range r {
			switch ri {
			case -1:
				for i := range xs {
					xs[i] += "a"
				}
			case -2:
				for i := range xs {
					xs[i] += "b"
				}
			default:
				results := traverse(rules, ri)
				newXs := []string{}
				for _, result := range results {
					for _, x := range xs {
						newXs = append(newXs, x+result)
					}
				}
				xs = newXs
			}
		}
		results = append(results, xs...)
	}
	return results
}

var rules = [][][]int{
	{{8, 11}},
	{{82, 39}, {119, 20}},
	{{39, 66}, {20, 115}},
	{{20, 128}, {39, 110}},
	{{12, 38}},
	{{20, 88}, {39, 27}},
	{{20, 10}, {39, 18}},
	{{5, 39}, {111, 20}},
	{{42}},
	{{20, 2}, {39, 74}},
	{{39, 20}, {39, 39}},
	{{42, 31}},
	{{20, 39}, {20, 20}},
	{{39, 132}, {20, 32}},
	{{39, 100}, {20, 114}},
	{{39, 9}, {20, 59}},
	{{39, 33}, {20, 46}},
	{{39, 10}, {20, 33}},
	{{20, 39}, {39, 39}},
	{{20, 113}, {39, 116}},
	{{-2}}, // This is "b"
	{{7, 39}, {105, 20}},
	{{39, 82}, {20, 40}},
	{{67, 39}, {129, 20}},
	{{40, 20}, {10, 39}},
	{{84, 39}, {112, 20}},
	{{20, 48}, {39, 83}},
	{{109, 20}, {40, 39}},
	{{130, 20}, {119, 39}},
	{{20, 16}, {39, 78}},
	{{20, 122}, {39, 52}},
	{{39, 43}, {20, 118}},
	{{39, 90}, {20, 57}},
	{{39, 39}},
	{{107, 39}, {10, 20}},
	{{39, 55}, {20, 4}},
	{{20, 107}, {39, 40}},
	{{106, 39}, {85, 20}},
	{{20}, {39}},
	{{-1}}, // This is "a"
	{{39, 20}},
	{{23, 39}, {37, 20}},
	{{20, 51}, {39, 120}},
	{{21, 39}, {103, 20}},
	{{79, 20}, {50, 39}},
	{{39, 73}, {20, 68}},
	{{20, 20}},
	{{82, 20}, {12, 39}},
	{{82, 39}, {12, 20}},
	{{39, 60}, {20, 88}},
	{{39, 36}, {20, 6}},
	{{39, 91}, {20, 117}},
	{{109, 39}, {10, 20}},
	{{94, 20}, {133, 39}},
	{{82, 38}},
	{{109, 38}},
	{{12, 39}, {109, 20}},
	{{107, 39}, {130, 20}},
	{{44, 39}, {77, 20}},
	{{95, 20}, {49, 39}},
	{{20, 128}, {39, 82}},
	{{20, 12}, {39, 33}},
	{{72, 20}, {28, 39}},
	{{20, 107}, {39, 70}},
	{{20, 96}, {39, 107}},
	{{25, 20}, {127, 39}},
	{{39, 128}, {20, 12}},
	{{92, 20}, {24, 39}},
	{{39, 119}, {20, 128}},
	{{20, 53}, {39, 14}},
	{{20, 39}, {39, 38}},
	{{39, 6}, {20, 92}},
	{{96, 39}, {110, 20}},
	{{96, 20}, {18, 39}},
	{{39, 76}, {20, 97}},
	{{39, 131}, {20, 35}},
	{{20, 10}, {39, 107}},
	{{98, 39}, {30, 20}},
	{{39, 130}, {20, 40}},
	{{39, 86}, {20, 36}},
	{{20, 10}, {39, 12}},
	{{125, 39}, {61, 20}},
	{{39, 39}, {20, 20}},
	{{96, 39}, {107, 20}},
	{{1, 39}, {99, 20}},
	{{108, 20}, {80, 39}},
	{{39, 96}, {20, 96}},
	{{20, 128}, {39, 12}},
	{{38, 18}},
	{{128, 20}, {107, 39}},
	{{20, 18}, {39, 96}},
	{{20, 13}, {39, 93}},
	{{12, 20}, {128, 39}},
	{{62, 20}, {71, 39}},
	{{20, 46}, {39, 109}},
	{{20, 102}, {39, 52}},
	{{20, 39}},
	{{40, 39}, {96, 20}},
	{{97, 20}, {56, 39}},
	{{46, 20}, {107, 39}},
	{{96, 20}, {109, 39}},
	{{114, 20}, {89, 39}},
	{{18, 20}, {82, 39}},
	{{39, 75}, {20, 19}},
	{{20, 18}, {39, 130}},
	{{39, 29}, {20, 101}},
	{{39, 3}, {20, 123}},
	{{20, 39}, {39, 20}},
	{{39, 107}, {20, 128}},
	{{39, 39}, {38, 20}},
	{{38, 38}},
	{{17, 39}, {48, 20}},
	{{22, 20}, {54, 39}},
	{{66, 20}, {47, 39}},
	{{39, 40}, {20, 18}},
	{{20, 18}, {39, 110}},
	{{108, 39}, {99, 20}},
	{{20, 126}, {39, 69}},
	{{15, 20}, {65, 39}},
	{{20, 38}, {39, 39}},
	{{41, 39}, {58, 20}},
	{{20, 64}, {39, 92}},
	{{39, 70}, {20, 96}},
	{{39, 10}, {20, 96}},
	{{20, 70}, {39, 33}},
	{{39, 96}, {20, 10}},
	{{20, 26}, {39, 45}},
	{{121, 39}, {81, 20}},
	{{39, 20}, {20, 38}},
	{{87, 39}, {34, 20}},
	{{39, 20}, {20, 20}},
	{{39, 63}, {20, 124}},
	{{39, 64}, {20, 104}},
	{{107, 20}, {12, 39}},
}
