package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// VALUETOGET is the value to add up to
var VALUETOGET = 2020

// ChallengeTwo solves the first challenge
func ChallengeTwo() {
	nums := make(map[int]bool)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		nums[val] = true
	}

	for i := range nums {
		for j := range nums {

			third := i + j

			if third > VALUETOGET {
				continue
			}

			tmp := VALUETOGET - third

			if _, ok := nums[tmp]; ok {
				ans := tmp * i * j

				// ChallengeTwo Answer
				fmt.Println(strconv.Itoa(ans))
				return
			}

		}
	}

}

// ChallengeOne solves the first challenge
func ChallengeOne() {

	nums := make(map[int]bool)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		nums[val] = true

		search := VALUETOGET - val

		if _, ok := nums[search]; ok {
			ans := search * val

			// ChallengeOne Answer
			fmt.Println(strconv.Itoa(ans))
			return
		}
	}

}

func main() {

	// ChallengeOne()
	ChallengeTwo()

}
