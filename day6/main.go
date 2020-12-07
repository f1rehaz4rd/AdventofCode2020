package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "input.txt"

func intersection(set1, set2 map[byte]bool) map[byte]bool {

	newset := make(map[byte]bool)

	for k1, _ := range set1 {
		if _, ok := set2[k1]; ok {
			newset[k1] = true
		}
	}

	return newset
}

func EvaluateAnswersTwo(lines []string) int {

	set := make(map[byte]bool)

	for i := 0; i < len(lines[0]); i++ {
		set[lines[0][i]] = true
	}

	if len(lines) > 1 {
		for i := 1; i < len(lines); i++ {
			tmpset := make(map[byte]bool)

			for j := 0; j < len(lines[i]); j++ {
				tmpset[lines[i][j]] = true
			}

			set = intersection(set, tmpset)

		}
	}

	return len(set)
}

func ChallengeTwo() {

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	holdStr := make([]string, 0)
	ans := 0

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) == 0 {
			ans += EvaluateAnswersTwo(holdStr)
			holdStr = make([]string, 0)
			continue
		}

		holdStr = append(holdStr, line)
	}

	ans += EvaluateAnswersTwo(holdStr)

	fmt.Printf("ans: %d\n", ans)

}

func EvaluateAnswers(str string) int {

	set := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		set[str[i]] = true
	}

	return len(set)
}

func ChallengeOne() {

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	holdStr := ""
	ans := 0

	for scanner.Scan() {

		line := strings.Trim(scanner.Text(), "\n")

		if len(line) == 0 {
			ans += EvaluateAnswers(holdStr)
			holdStr = ""
		}

		holdStr += line
	}

	ans += EvaluateAnswers(holdStr)
	holdStr = ""

	fmt.Printf("ans: %d\n", ans)

}

func main() {
	// ChallengeOne()
	ChallengeTwo()
}
