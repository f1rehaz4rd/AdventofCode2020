package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const FILENAME = "input.txt"
const TREECHAR = "#"
const RIGHTMOVEVAL = 3

func buildBoard(inputFile string) ([][]string, int, int) {

	board := [][]string{}

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	row := 0
	col := 0
	for scanner.Scan() {
		colStr := scanner.Text()

		tmpRow := []string{}
		for col = 0; col < len(colStr); col++ {
			tmpRow = append(tmpRow, string(colStr[col]))
		}

		board = append(board, tmpRow)
		row++
	}

	return board, row, col
}

func ChallengeOne() {

	board, row, col := buildBoard(FILENAME)

	trees := 0
	j := RIGHTMOVEVAL
	for i := 1; i < row; i++ {
		if board[i][j] == TREECHAR {
			trees++
		}

		j = (j + RIGHTMOVEVAL) % col

	}

	// Print Answer
	fmt.Printf("# number of trees: %d\n", trees)

}

func main() {
	ChallengeOne()
}
