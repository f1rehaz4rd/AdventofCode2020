package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const FILENAME = "input.txt"

func ChallengeTwoValidate(passport map[string]string) bool {

	valid := false

	if len(passport) > 8 {
		return valid
	}

	fieldCheck := map[string]func(string) bool{
		"byr": func(str string) bool {
			num, _ := strconv.Atoi(str)
			return num >= 1920 && num <= 2002
		},
		"iyr": func(str string) bool {
			num, _ := strconv.Atoi(str)
			return num >= 2010 && num <= 2020
		},
		"eyr": func(str string) bool {
			num, _ := strconv.Atoi(str)
			return num >= 2020 && num <= 2030
		},

		"hgt": func(str string) bool {
			if str[len(str)-2:len(str)] == "cm" {
				num, _ := strconv.Atoi(str[:len(str)-2])
				return num >= 150 && num <= 193
			} else if str[len(str)-2:len(str)] == "in" {
				num, _ := strconv.Atoi(str[:len(str)-2])
				return num >= 59 && num <= 76
			}
			return false
		},
		"hcl": func(str string) bool {
			regexCheck := "^#[0-9a-f]{6}$"
			if res, _ := regexp.MatchString(regexCheck, str); res {
				return true
			}
			return false
		},
		"ecl": func(str string) bool {
			regexCheck := "^(amb|blu|brn|gry|grn|hzl|oth)"
			if res, _ := regexp.MatchString(regexCheck, str); res {
				return true
			}
			return false
		},
		"pid": func(str string) bool {
			regexCheck := "^\\d{9}$"
			if res, _ := regexp.MatchString(regexCheck, str); res {
				return true
			}
			return false
		},
		"cid": func(str string) bool {
			return false
		},
	}

	check := 0
	for key, val := range passport {

		if fieldCheck[key](val) {
			check++
		}

		if check == 7 {
			valid = true
		}

	}

	return valid
}

func ChallengeTwo() {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	passport := make(map[string]string)
	numOfValid := 0

	for scanner.Scan() {

		// line := strings.Trim(scanner.Text(), "\n")
		line := scanner.Text()

		// Check if there is a new passport
		if len(line) == 0 {

			//
			// Check if the passport is valid
			//
			if ChallengeTwoValidate(passport) {
				numOfValid++
			}

			// Clear passport and continue
			passport = make(map[string]string)
			continue
		}

		// Parse the passport lines
		spaceSplit := strings.Split(line, " ")
		for i := 0; i < len(spaceSplit); i++ {
			entry := strings.Split(spaceSplit[i], ":")

			passport[entry[0]] = entry[1]
		}

	}

	//
	// Checks after the file is EOF
	//
	if ChallengeTwoValidate(passport) {
		numOfValid++
	}

	fmt.Printf("# of valid: %d\n", numOfValid)
}

func ChallengeOneValidate(passport map[string]string) bool {

	valid := false
	fieldCheck := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}

	check := 0
	for key, _ := range fieldCheck {
		if _, ok := passport[key]; ok {
			fieldCheck[key] = true
			check++
			if check == 7 {
				valid = true
			}
		}
	}

	return valid
}

func ChallengeOne() {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	passport := make(map[string]string)
	numOfValid := 0

	for scanner.Scan() {

		// line := strings.Trim(scanner.Text(), "\n")
		line := scanner.Text()

		// Check if there is a new passport
		if len(line) == 0 {

			//
			// Check if the passport is valid
			//
			if ChallengeOneValidate(passport) {
				numOfValid++
			}

			// Clear passport and continue
			passport = make(map[string]string)
			continue
		}

		// Parse the passport lines
		spaceSplit := strings.Split(line, " ")
		for i := 0; i < len(spaceSplit); i++ {
			entry := strings.Split(spaceSplit[i], ":")

			passport[entry[0]] = entry[1]
		}

	}

	//
	// Checks after the file is EOF
	//
	if ChallengeOneValidate(passport) {
		numOfValid++
	}

	fmt.Printf("# of valid: %d\n", numOfValid)
}

func main() {
	// ChallengeOne()
	ChallengeTwo()
}
