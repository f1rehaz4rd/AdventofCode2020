package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ChallengeOneValidate validates based of the first challenge's
// parameters for the password policy
func ChallengeOneValidate(entry string) bool {

	valid := false

	//
	// Parse the entry to get values needed
	//
	tmpSplit := strings.Split(entry, ":")

	minVal, _ := strconv.Atoi(strings.Split(tmpSplit[0], "-")[0])
	maxVal, _ := strconv.Atoi(strings.Split(strings.Split(tmpSplit[0], "-")[1], " ")[0])

	letter := strings.Split(tmpSplit[0], " ")[1]
	password := strings.TrimSpace(tmpSplit[1])

	//
	// Do validation
	//
	letterCount := strings.Count(password, letter)

	if letterCount >= minVal && letterCount <= maxVal {
		valid = true
	}

	// Return if it was valid
	return valid
}

// ChallengeOne Solves and prints the answer to part 1
func ChallengeOne() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ctr := 0

	for scanner.Scan() {
		if ChallengeOneValidate(scanner.Text()) {
			ctr++
		}
	}

	fmt.Printf("# of valid passwords: %d\n", ctr)
}

func main() {
	ChallengeOne()
}
