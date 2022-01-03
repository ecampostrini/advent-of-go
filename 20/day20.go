package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"github.com/ecampostrini/advent-of-go/utils/types"
)

func readImage(scanner *bufio.Scanner) []string {
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func pixels2Num(in string) int {
	if len(in) > 9 {
		panic("pixels2Num requires the input to be of length <= 9")
	}
	var ret int
	for i := 0; i < len(in); i++ {
		var bit int
		if in[i] == '#' {
			bit = 1
		}
		ret = ret<<1 | bit
	}
	return ret
}

const directions = []types.Point{
	Point{-1, -1},
	Point{0, -1},
	Point{1, -1},
	Point{-1, 0},
	Point{0, 0},
	Point{0, 1},
	Point{-1, 1},
	Point{0, 1},
	Point{1, 1},
}

func getPixelSurroundings(x, y int, image []string) string {
	var ret string
	for i := 0; i < len(directions); i++ {
		direction := directions[i]
		iy := y + direction.Y
		ix := x + direction.X
		if iy < 0 || iy >= len(image) || ix < 0 || ix >= len(image[0]) {
			ret = ret + "."
		} else {
			ret = ret + image[iy][ix]
		}
	}
	return ret
}

func enhaceImage(image []string, algorithm string) []string {
	var ret []string = make([]string, len(image)+2)
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[i]); x++ {
			pixelSurroundings := getPixelSurroundings(x, y, image)
			ret[y] = ret[y] + algorithm[pixels2Num(pixelSurroundings)]
		}
	}
	return ret
}

func main() {
	scanner, file := files.ReadFile("./input.test.txt")
	defer file.Close()

	scanner.Scan()
	algorithm := scanner.Text()

	inputImage := readImage(scanner)
	fmt.Printf("Algorithm: %d\n", len(algorithm))
	fmt.Printf("Image: %v\n", inputImage)
	fmt.Printf("Test: %d\n", pixels2Num("...#...#."))
}
