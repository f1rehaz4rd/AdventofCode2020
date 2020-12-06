package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FILENAME = "input.txt"

func ChallengeTwo() {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	list := make(map[int64]bool)
	highest := int64(0)

	for scanner.Scan() {

		str := scanner.Text()

		byteStr := ""
		for i := 0; i < len(str); i++ {
			if str[i] == 'F' || str[i] == 'L' {
				byteStr = byteStr + "0"
			} else {
				byteStr = byteStr + "1"
			}
		}

		seatID, _ := strconv.ParseInt(byteStr, 2, 64)

		list[seatID] = true

		if seatID > highest {
			highest = seatID
		}
	}

	for i := int64(0); i < highest; i++ {
		if _, ok := list[i]; !ok {
			if _, ok := list[i-1]; ok {
				if _, ok := list[i+1]; ok {
					// Print the answer
					fmt.Printf("Highest seat ID: %d\n", i)
					break
				}
			}
		}
	}

}

func ChallengeOne() {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	highest := int64(0)

	for scanner.Scan() {

		str := scanner.Text()

		byteStr := ""
		for i := 0; i < len(str); i++ {
			if str[i] == 'F' || str[i] == 'L' {
				byteStr = byteStr + "0"
			} else {
				byteStr = byteStr + "1"
			}
		}

		seatID, _ := strconv.ParseInt(byteStr, 2, 64)

		if seatID > highest {
			highest = seatID
		}

	}

	// Print the ans
	fmt.Printf("Highest seat ID: %d\n", highest)
}

func main() {
	// ChallengeOne()
	ChallengeTwo()
}
