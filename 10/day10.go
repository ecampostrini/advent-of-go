package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var pentaltyPoints = map[int]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionPoints = map[int]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
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

func getFirstMismatch(line string) (int, []int) {
	stack, stackPos := make([]int, len(line)), 0
	for i := 0; i < len(line); i++ {
		currentChar := int(line[i])
		if pentaltyPoints[currentChar] == 0 {
			stack[stackPos] = currentChar
			stackPos++
		} else {
			if !areAMatch(stack[stackPos-1], currentChar) {
				return currentChar, stack
			}
			stackPos--
		}
	}
	return 0, stack[:stackPos]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	syntaxErrorScore := 0
	var completionScoreList []int
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		firstMismatch, stack := getFirstMismatch(scanner.Text())
		if firstMismatch == 0 {
			completionScore := 0
			for i := range stack {
				currentChar := stack[len(stack)-i-1]
				completionScore = (completionScore * 5) + completionPoints[currentChar]
			}
			completionScoreList = append(completionScoreList, completionScore)
		} else {
			syntaxErrorScore += pentaltyPoints[firstMismatch]
		}
	}

	// part 1
	fmt.Println(syntaxErrorScore)

	// part 2
	sort.Ints(completionScoreList)
	fmt.Println(completionScoreList[len(completionScoreList)/2])
}
