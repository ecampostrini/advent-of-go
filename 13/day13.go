package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeBidimensionalSliceString(dx, dy int) [][]string {
	ret := make([][]string, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]string, dx)
	}
	return ret
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, c := range row {
			if c == "" {
				fmt.Printf(".")
			} else {
				fmt.Printf("%s", c)
			}
		}
		fmt.Printf("\n")
	}
}

func readGrid(scanner *bufio.Scanner) [][]string {
	points := make(map[int][]int)
	var maxX, maxY int

	scanner.Scan()
	for scanner.Text() != "" {
		point := strings.Split(scanner.Text(), ",")
		var x, y int
		var err error

		if x, err = strconv.Atoi(point[0]); err != nil {
			panic(err)
		}

		if y, err = strconv.Atoi(point[1]); err != nil {
			panic(err)
		}

		points[y] = append(points[y], x)

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		scanner.Scan()
	}

	ret := makeBidimensionalSliceString(maxX+1, maxY+1)
	for i := 0; i <= maxY; i++ {
		for _, x := range points[i] {
			ret[i][x] = "#"
		}
	}
	return ret
}

type Point struct {
	X, Y int
}

func readFolds(scanner *bufio.Scanner) []Point {
	var ret []Point
	for scanner.Scan() {
		line := strings.Split(strings.TrimPrefix(scanner.Text(), "fold along "), "=")

		var pos int
		var err error
		if pos, err = strconv.Atoi(line[1]); err != nil {
			panic(err)
		}

		if line[0] == "x" {
			ret = append(ret, Point{pos, 0})
		} else {
			ret = append(ret, Point{0, pos})
		}
	}
	return ret
}

func main() {
	file, err := os.Open("./input.test.txt")
	if err != nil {
		fmt.Println("Failed to read input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := readGrid(scanner)
	printGrid(grid)

	folds := readFolds(scanner)
	fmt.Printf("%v\n", folds)
}
