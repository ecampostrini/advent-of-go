package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

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

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	heightmap := makeHeightmap(bufio.NewScanner(file))
	totalRiskLevel := 0
	for y, row := range heightmap {
		for x, currHeight := range row {
			if isLowPoint(x, y, heightmap) {
				totalRiskLevel += 1 + currHeight
			}
		}
	}

	fmt.Println(totalRiskLevel)

}
