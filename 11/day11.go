package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X, Y int
}

func readGrid(scanner *bufio.Scanner) [][]int {
	var ret [][]int
	for scanner.Scan() {
		line := scanner.Text()
		gridRow := make([]int, len(line))
		for i, c := range scanner.Text() {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println("Error while parsing input: ", err)
				os.Exit(1)
			}
			gridRow[i] = n
		}
		ret = append(ret, gridRow)
	}
	return ret
}

func printGrid(grid [][]int) {
	fmt.Println("---")
	for _, row := range grid {
		fmt.Println(row)
	}
}

var directions = []Point{
	Point{0, 1},
	Point{1, 1},
	Point{1, 0},
	Point{1, -1},
	Point{0, -1},
	Point{-1, -1},
	Point{-1, 0},
	Point{-1, 1},
}

func isWithinGrid(p Point, grid [][]int) bool {
	if p.X < 0 || p.X >= len(grid[0]) {
		return false
	}
	if p.Y < 0 || p.Y >= len(grid) {
		return false
	}
	return true
}

func propagate(grid [][]int, flashQueue []Point) int {
	ret := 0
	for len(flashQueue) > 0 {
		current := flashQueue[0]
		flashQueue = flashQueue[1:]
		for _, direction := range directions {
			next := Point{current.X + direction.X, current.Y + direction.Y}
			if isWithinGrid(next, grid) && grid[next.Y][next.X] > 0 && grid[next.Y][next.X] <= 9 {
				grid[next.Y][next.X]++
				if grid[next.Y][next.X] > 9 {
					flashQueue = append(flashQueue, next)
				}
			}
		}
		grid[current.Y][current.X] = 0
		ret++
	}
	return ret
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error while opening input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	grid := readGrid(bufio.NewScanner(file))
	totalFlashCount := 0
	for i := 0; i < 500; i++ {
		var flashQueue []Point
		for y, row := range grid {
			for x, _ := range row {
				grid[y][x]++
				if grid[y][x] > 9 {
					flashQueue = append(flashQueue, Point{x, y})
				}
			}
		}
		currentFlashCount := propagate(grid, flashQueue)
		totalFlashCount += currentFlashCount

		// part 1
		if i == 100 {
			fmt.Println(totalFlashCount)
		}

		// part 2
		if currentFlashCount == len(grid)*len(grid[0]) {
			fmt.Println(i + 1)
			break
		}
	}
}
