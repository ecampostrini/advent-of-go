package main

import (
	"bufio"
	"fmt"
	"os"
)

var points = map[int]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var pairs = map[int]int{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func areAMatch(c1, c2 int) bool {
	if pairs[c1] == c2 {
		return true
	}
	return false
}

func getFirstMismatch(line string) int {
	stack, stackPos := make([]int, len(line)), 0
	for i := 0; i < len(line); i++ {
		currentChar := int(line[i])
		if points[currentChar] == 0 {
			stack[stackPos] = currentChar
			stackPos++
		} else {
			if !areAMatch(stack[stackPos-1], currentChar) {
				return currentChar
			}
			stackPos--
		}
	}
	return 0
}

func main() {
	file, err := os.Open("./input.test.txt")
	if err != nil {
		fmt.Println("Failed to open input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	syntaxErroScore := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		syntaxErroScore += points[getFirstMismatch(scanner.Text())]
	}
	fmt.Println(syntaxErroScore)
}
