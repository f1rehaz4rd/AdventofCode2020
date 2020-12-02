package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

		search := 2020 - val

		if _, ok := nums[search]; ok {
			ans := search * val
			strans := strconv.Itoa(ans)

			fmt.Println(strans)
		}
	}

}

func main() {

	ChallengeOne()

}
