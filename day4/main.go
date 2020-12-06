package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "input.txt"

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
	ChallengeOne()
}
