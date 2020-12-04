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

var invalid = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`

var valid = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

func TestParseInput(t *testing.T) {
	passports := parseInput(input)
	assert.Equal(t, 4, len(passports))
	assert.Equal(t, 8, len(passports[0]))
	assert.Equal(t, 7, len(passports[1]))
	assert.Equal(t, 7, len(passports[2]))
	assert.Equal(t, 6, len(passports[3]))
}

func TestInvalid(t *testing.T) {
	for _, passport := range parseInput(invalid) {
		assert.False(t, validate(passport))
	}
}

func TestValid(t *testing.T) {
	for _, passport := range parseInput(valid) {
		assert.True(t, validate(passport))
	}
}

func TestFieldValidation(t *testing.T) {

	// byr valid:   2002
	// byr invalid: 2003
	assert.True(t, validateBYR("2002"))
	assert.False(t, validateBYR("2003"))
	assert.True(t, validateBYR("1920"))
	assert.False(t, validateBYR("1919"))

	assert.True(t, validateIYR("2020"))
	assert.False(t, validateIYR("2021"))
	assert.True(t, validateIYR("2010"))
	assert.False(t, validateIYR("2009"))

	assert.True(t, validateEYR("2030"))
	assert.False(t, validateEYR("2031"))
	assert.True(t, validateEYR("2020"))
	assert.False(t, validateEYR("2019"))

	// hgt valid:   60in
	// hgt valid:   190cm
	// hgt invalid: 190in
	// hgt invalid: 190
	assert.True(t, validateHGT("60in"))
	assert.True(t, validateHGT("190cm"))
	assert.False(t, validateHGT("190in"))
	assert.False(t, validateHGT("190"))

	// hcl valid:   #123abc
	// hcl invalid: #123abz
	// hcl invalid: 123abc
	assert.True(t, validateHCL("#123abc"))
	assert.False(t, validateHCL("#123abz"))
	assert.False(t, validateHCL("123abc"))

	// ecl valid:   brn
	// ecl invalid: wat
	assert.True(t, validateECL("brn"))
	assert.False(t, validateECL("wat"))

	// pid valid:   000000001
	// pid invalid: 0123456789
	assert.True(t, validatePID("000000001"))
	assert.False(t, validatePID("0123456789"))

}

func TestCountValid(t *testing.T) {
	passports := parseInput(input)
	assert.Equal(t, 2, countValid(passports))
}
