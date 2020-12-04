package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestParseInput(t *testing.T) {
	passports := parseInput(input)
	assert.Equal(t, 4, len(passports))
	assert.Equal(t, 8, len(passports[0]))
	assert.Equal(t, 7, len(passports[1]))
	assert.Equal(t, 7, len(passports[2]))
	assert.Equal(t, 6, len(passports[3]))
}

func TestValidate(t *testing.T) {
	passports := parseInput(input)
	assert.True(t, validate(passports[0]))
	assert.False(t, validate(passports[1]))
	assert.True(t, validate(passports[2]))
	assert.False(t, validate(passports[3]))
}

func TestCountValid(t *testing.T) {
	passports := parseInput(input)
	assert.Equal(t, 2, countValid(passports))
}
