package main

import (
	"fmt"
	"strings"

	"github.com/scottyw/adventofcode2020/pkg/aoc"
)

func main() {
	passportInput := aoc.FileToString("input/04.txt")
	fmt.Println(countValid(parseInput(passportInput)))
}

func parseInput(input string) []map[string]string {

	// We know the input is validly formatted and that
	// a single blank line is used to separate records
	// so let's replace double newlines with our own
	// separator for now
	input = strings.ReplaceAll(input, "\n\n", "||")

	// Any remaining newlines are in the middle of
	// records so let's remove those so that there
	// are no newlines in the input at all
	input = strings.ReplaceAll(input, "\n", " ")

	// Use the || separator we added at the start to separate the records
	records := strings.Split(input, "||")

	// Convert each record into a map
	passports := []map[string]string{}
	for _, record := range records {
		passport := map[string]string{}
		fields := strings.Split(record, " ")
		for _, field := range fields {
			// Empty fields are OK, they're just a consequence of splitting
			// with leading or trailing spaces
			if field == "" {
				continue
			}
			values := strings.Split(field, ":")
			if len(values) != 2 {
				panic(fmt.Sprintf("Malformed field '%s' in record '%s'", field, record))
			}
			passport[values[0]] = values[1]
		}
		passports = append(passports, passport)
	}

	return passports
}

func validate(passport map[string]string) bool {
	// Ignore extra fields we don't care about and
	// just check for the 7 fields we care about
	return passport["byr"] != "" &&
		passport["hgt"] != "" &&
		passport["hcl"] != "" &&
		passport["iyr"] != "" &&
		passport["eyr"] != "" &&
		passport["ecl"] != "" &&
		passport["pid"] != ""
}

func countValid(passports []map[string]string) int {
	var count int
	for _, passport := range passports {
		if validate(passport) {
			count++
		}
	}
	return count
}
