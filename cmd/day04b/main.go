package main

import (
	"fmt"
	"regexp"
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
	return validateBYR(passport["byr"]) &&
		validateHGT(passport["hgt"]) &&
		validateHCL(passport["hcl"]) &&
		validateIYR(passport["iyr"]) &&
		validateEYR(passport["eyr"]) &&
		validateECL(passport["ecl"]) &&
		validatePID(passport["pid"])
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func validateBYR(v string) bool {
	return validateYR(v, 1920, 2002)
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func validateIYR(v string) bool {
	return validateYR(v, 2010, 2020)
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func validateEYR(v string) bool {
	return validateYR(v, 2020, 2030)
}

func validateYR(v string, min, max int) bool {
	if v == "" {
		return false
	}
	y := aoc.Atoi(v)
	return y >= min && y <= max
}

// hgt (Height) - a number followed by either cm or in:
// 	If cm, the number must be at least 150 and at most 193.
// 	If in, the number must be at least 59 and at most 76.
var hgt = regexp.MustCompile("(\\d+)(cm|in)")

func validateHGT(v string) bool {
	matches := hgt.FindStringSubmatch(v)
	if matches == nil {
		return false
	}
	h := aoc.Atoi(matches[1])
	u := matches[2]
	switch u {
	case "in":
		return h >= 59 && h <= 76
	case "cm":
		return h >= 150 && h <= 193
	default:
		return false
	}
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
var hcl = regexp.MustCompile("#[0-9a-f]{6}")

func validateHCL(v string) bool {
	return hcl.FindStringSubmatch(v) != nil
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validateECL(v string) bool {
	switch v {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
var pid = regexp.MustCompile("^[0-9]{9}$")

func validatePID(v string) bool {
	return pid.FindStringSubmatch(v) != nil
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
