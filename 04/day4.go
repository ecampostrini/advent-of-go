package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board = [][]string

func readBoard(scanner *bufio.Scanner) Board {
	var ret Board
	for scanner.Text() == "" {
		scanner.Scan()
	}
	for scanner.Text() != "" {
		var row []string
		for _, number := range strings.Split(scanner.Text(), " ") {
			if number == "" {
				continue
			}
			row = append(row, strings.TrimSpace(number))
		}
		ret = append(ret, row)
		scanner.Scan()
	}
	return ret
}

func markNumber(board *Board, number string) {
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			if (*board)[i][j] == number {
				(*board)[i][j] = "X"
			}
		}
	}
}

func hasFullCol(b *Board) bool {
	board := *b
	flag := false
	for i := 0; !flag && i < len(board[0]); i++ {
		flag = true
		for j := 0; flag && j < len(board); j++ {
			flag = flag && (board[j][i] == "X")
		}
	}
	return flag
}

func hasFullRow(b *Board) bool {
	board := *b
	flag := false
	for i := 0; !flag && i < len(board); i++ {
		flag = true
		for j := 0; flag && j < len(board[i]); j++ {
			flag = flag && board[i][j] == "X"
		}
	}
	return flag
}

func sumNumbersInBoard(b *Board) int {
	board := *b
	ret := 0
	for _, row := range board {
		for _, num := range row {
			n, _ := strconv.Atoi(num)
			ret += n
		}
	}
	return ret
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Oops")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")

	var boards []Board
	hasMore := scanner.Scan()
	for hasMore {
		boards = append(boards, readBoard(scanner))
		hasMore = scanner.Scan()
	}

	// part 1
	var winnerBoard *Board
	var lastNumberDrawn string
	for i := 0; winnerBoard == nil && i < len(numbers); i++ {
		for j := 0; winnerBoard == nil && j < len(boards); j++ {
			lastNumberDrawn = numbers[i]
			board := &boards[j]
			markNumber(board, lastNumberDrawn)
			if hasFullRow(board) || hasFullCol(board) {
				winnerBoard = board
			}
		}
	}

	if winnerBoard == nil {
		fmt.Println("No winner found!")
		os.Exit(1)
	}

	lnd, _ := strconv.Atoi(lastNumberDrawn)
	fmt.Println(sumNumbersInBoard(winnerBoard) * lnd)

	// part 2
	winMap := make(map[int]bool)
	for i := 0; len(winMap) < len(boards) && i < len(numbers); i++ {
		for j := 0; len(winMap) < len(boards) && j < len(boards); j++ {
			if winMap[j] {
				continue
			}
			lastNumberDrawn = numbers[i]
			board := &boards[j]
			markNumber(board, lastNumberDrawn)
			if hasFullRow(board) || hasFullCol(board) {
				winnerBoard = board
				winMap[j] = true
			}
		}
	}

	lnd, _ = strconv.Atoi(lastNumberDrawn)
	fmt.Println(sumNumbersInBoard(winnerBoard) * lnd)
}
