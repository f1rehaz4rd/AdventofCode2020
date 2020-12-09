package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "input.txt"
const BAGSPLIT = " contain "
const CONTAINSNONE = "no other bags"
const MYBAG = "shiny gold bag"

func checkBag(start string, dict map[string][]string) bool {

	if start == MYBAG {
		return true
	} else if start == CONTAINSNONE {
		return false
	}

	otherBags := dict[start]
	for i := 0; i < len(otherBags); i++ {
		if checkBag(otherBags[i], dict) {
			return true
		}
	}

	return false
}

func ChallengeOne() {

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	dict := make(map[string][]string)

	//
	// Parse all the rules and make a dict
	//
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), BAGSPLIT)

		rules := strings.Split(line[1], ",")

		tmp := make([]string, 0)
		for i := 0; i < len(rules); i++ {
			rule := strings.Trim(rules[i], " .")
			if rule == CONTAINSNONE {
				tmp = []string{rule}
				break
			} else {
				// Takes the 's' out of the bag
				if rule[len(rule)-1] == 's' {
					tmp = append(tmp, rule[2:len(rule)-1])
				} else {
					tmp = append(tmp, rule[2:])
				}
			}
		}

		dict[line[0][:len(line[0])-1]] = tmp
	}

	//
	// Go through the rules and check for the gold bag
	//
	ans := 0
	for k, _ := range dict {
		if k == MYBAG {
			continue
		} else if checkBag(k, dict) {
			ans++
		}
	}

	fmt.Printf("ans: %d\n", ans)
}

// func main() {
// 	ChallengeOne()
// }
