package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	for _, row := range grid {
		fmt.Println(row)
	}
}

func main() {
	file, err := os.Open("./input.test.txt")
	if err != nil {
		fmt.Println("Error while opening input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	grid := readGrid(bufio.NewScanner(file))
	printGrid(grid)
}
