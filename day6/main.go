package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "input.txt"

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
	ChallengeOne()
}
