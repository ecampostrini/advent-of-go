package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"github.com/ecampostrini/advent-of-go/utils/slices"
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

var directions = []types.Point{
	types.Point{-1, -1},
	types.Point{0, -1},
	types.Point{1, -1},
	types.Point{-1, 0},
	types.Point{0, 0},
	types.Point{1, 0},
	types.Point{-1, 1},
	types.Point{0, 1},
	types.Point{1, 1},
}

func getPixelSurroundings(x, y int, image []string) string {
	var ret string
	for i := 0; i < len(directions); i++ {
		direction := directions[i]
		iy := y + direction.Y
		ix := x + direction.X
		if iy < 0 || iy >= len(image) || ix < 0 || ix >= len(image[0]) {
			ret = ret + "."
			//fmt.Printf("(%d, %d): %s\n", ix, iy, ".")
		} else {
			ret = ret + string(image[iy][ix])
			//fmt.Printf("(%d, %d): %s\n", ix, iy, string(image[iy][ix]))
		}
	}
	return ret
}

func enhaceImage(image []string, algorithm string) []string {
	var ret []string = make([]string, len(image)+2)
	for y := -1; y < len(image)+1; y++ {
		for x := -1; x < len(image[0])+1; x++ {
			pixelSurroundings := getPixelSurroundings(x, y, image)
			ret[y+1] = ret[y+1] + string(algorithm[pixels2Num(pixelSurroundings)])
		}
	}
	return ret
}

func countPixels(image []string, pixel rune) int {
	var ret int
	for _, line := range image {
		for _, c := range line {
			if c == pixel {
				ret++
			}
		}
	}
	return ret
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	scanner.Scan()
	algorithm := scanner.Text()
	fmt.Printf("Algorithm: %d\n", len(algorithm))
	// consume empty space
	scanner.Scan()
	inputImage := readImage(scanner)
	var outputImage []string = inputImage
	for i := 0; i < 1; i++ {
		outputImage = enhaceImage(outputImage, algorithm)
		fmt.Printf("---\n")
		slices.PrintStringSlice(outputImage)
		fmt.Printf("\nPixels lit: %d\n", countPixels(outputImage, '#'))
	}
}
