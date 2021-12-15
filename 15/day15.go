package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
  //"github.com/ecampostrini/advent-of-go/utils/slices"
	"math"
)

func ParseGridInt(scanner *bufio.Scanner) [][]int {
	var ret [][]int
	for scanner.Scan() {
		textLine := scanner.Text()
		newRow := make([]int, len(textLine))
		for i, c := range textLine {
			newRow[i] = int(c - '0')
		}
		ret = append(ret, newRow)
	}
	return ret
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	grid := ParseGridInt(scanner)
  //slices.PrintGridInt(grid)

	grid[0][0] = 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if i == 0 && j == 0 {
				continue
			}
			upVal, leftVal := math.MaxInt, math.MaxInt
			if i > 0 {
				upVal = grid[i-1][j]
			}
			if j > 0 {
				leftVal = grid[i][j-1]
			}
			grid[i][j] += Min(upVal, leftVal)
		}
	}

  //fmt.Println("---")
  //slices.PrintGridInt(grid)
	fmt.Println(grid[len(grid)-1][len(grid[0])-1])
}
