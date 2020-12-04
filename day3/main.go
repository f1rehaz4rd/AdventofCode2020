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

var RIGHTMOVEVALS = []int{1, 3, 5, 7}

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

func ChallengeTwo() {
	board, row, col := buildBoard(FILENAME)

	ans := 1
	trees := []int{0, 0, 0, 0, 0}

	rightMoves := []int{1, 3, 5, 7}
	downTwoCheck := 1

	for i := 1; i < row; i++ {
		//
		// Check the single down moves
		//
		for j := 0; j < len(rightMoves); j++ {
			if board[i][rightMoves[j]] == TREECHAR {
				trees[j]++
			}

			rightMoves[j] = (rightMoves[j] + RIGHTMOVEVALS[j]) % col
		}

		//
		// Check the down 2
		//
		if (i % 2) == 0 {
			if board[i][downTwoCheck] == TREECHAR {
				trees[4]++
			}

			downTwoCheck = (downTwoCheck + 1) % col
		}
	}

	//
	// Multiply them together
	//
	for i := 0; i < len(trees); i++ {
		ans = ans * trees[i]
	}

	// Print Answer
	fmt.Printf("ans: %d\n", ans)
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
	// ChallengeOne()
	ChallengeTwo()
}
