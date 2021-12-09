package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

func makeHeightmap(scanner *bufio.Scanner) [][]int {
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

func isLowPoint(x, y int, heightmap [][]int) bool {
	target := heightmap[y][x]
	if x-1 >= 0 && target >= heightmap[y][x-1] {
		return false
	}
	if x+1 < len(heightmap[0]) && target >= heightmap[y][x+1] {
		return false
	}
	if y-1 >= 0 && target >= heightmap[y-1][x] {
		return false
	}
	if y+1 < len(heightmap) && target >= heightmap[y+1][x] {
		return false
	}

	return true
}

func getBasinSize(x, y int, heightmap [][]int) int {
	var helper func(int, int)
	visited := make(map[Point]bool)
	count := 1

	basinCondition := func(x, y, target int) bool {
		if heightmap[y][x] > target && heightmap[y][x] < 9 && !visited[Point{x, y}] {
			return true
		}
		return false
	}

	helper = func(x, y int) {
		if x < 0 || x >= len(heightmap[0]) || y < 0 || y >= len(heightmap) {
			return
		}

		target := heightmap[y][x]
		if x-1 >= 0 && basinCondition(x-1, y, target) {
			count++
			visited[Point{x - 1, y}] = true
			helper(x-1, y)
		}
		if x+1 < len(heightmap[0]) && basinCondition(x+1, y, target) {
			count++
			visited[Point{x + 1, y}] = true
			helper(x+1, y)
		}
		if y-1 >= 0 && basinCondition(x, y-1, target) {
			count++
			visited[Point{x, y - 1}] = true
			helper(x, y-1)
		}
		if y+1 < len(heightmap) && basinCondition(x, y+1, target) {
			count++
			visited[Point{x, y + 1}] = true
			helper(x, y+1)
		}
	}
	helper(x, y)

	return count
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Failed to open input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	heightmap := makeHeightmap(bufio.NewScanner(file))
	totalRiskLevel, lowPointsCount := 0, 0
	for y, row := range heightmap {
		for x, currHeight := range row {
			if isLowPoint(x, y, heightmap) {
				lowPointsCount++
				totalRiskLevel += 1 + currHeight
			}
		}
	}
	fmt.Println(totalRiskLevel)

	var basinSizeList []int
	for y, row := range heightmap {
		for x, currHeight := range row {
			if isLowPoint(x, y, heightmap) {
				basinSizeList = append(basinSizeList, getBasinSize(x, y, heightmap))
				totalRiskLevel += 1 + currHeight
			}
		}
	}

	sort.Ints(basinSizeList)
	acc := 1
	for _, basinSize := range basinSizeList[len(basinSizeList)-3:] {
		acc *= basinSize
	}
	fmt.Println(acc)
}
