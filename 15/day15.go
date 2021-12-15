package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"github.com/ecampostrini/advent-of-go/utils/slices"
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

func doit(grid [][]int) int {
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
  return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	grid := ParseGridInt(scanner)
	//slices.PrintGridInt(grid)

	const dimension = 5
	dy, dx := len(grid), len(grid[0])
	fmt.Printf("dy: %d, dx: %d\n", dy*dimension, dx*dimension)
	newGrid := make([][]int, dy*dimension)
	for i, _ := range newGrid {
		newGrid[i] = make([]int, dx*dimension)
	}

	for i, _ := range newGrid {
		for j, _ := range newGrid[i] {
			oi, oj := i%dy, j%dx
			offsetY, offsetX := i/dy, j/dx
      value := grid[oi][oj] + offsetY + offsetX
      newGrid[i][j] = (value / 10) + (value % 10)
		}
	}

	fmt.Println("\n---\n")
	slices.PrintGridInt(newGrid, "")
  fmt.Println(doit(newGrid))
	fmt.Println("\n---\n")
  slices.PrintGridInt(newGrid, " ")

	//grid[0][0] = 0
	//for i, _ := range grid {
	//for j, _ := range grid[i] {
	//if i == 0 && j == 0 {
	//continue
	//}
	//upVal, leftVal := math.MaxInt, math.MaxInt
	//if i > 0 {
	//upVal = solution[i-1][j]
	//}
	//if j > 0 {
	//leftVal = solution[i][j-1]
	//}
	//solution[i][j] += Min(upVal, leftVal)
	//}
	//}

	//fmt.Println("---")
	//slices.PrintGridInt(grid)
}
