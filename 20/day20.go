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
		ret = ret << 1
		if in[i] == '#' {
			ret |= 1
		}
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

func getPixelSurroundings(x, y int, image []string, darkVoid bool) string {
	var ret string
	for i := 0; i < len(directions); i++ {
		direction := directions[i]
		iy := y + direction.Y
		ix := x + direction.X
		if iy < 0 || iy >= len(image) || ix < 0 || ix >= len(image[0]) {
			if darkVoid {
				ret = ret + "."
			} else {
				ret = ret + "#"
			}
		} else {
			ret = ret + string(image[iy][ix])
		}
	}
	return ret
}

func enhaceImage(image []string, algorithm string, iterationNum int) []string {
	var ret []string = make([]string, len(image)+2)
	for y := -1; y < len(image)+1; y++ {
		for x := -1; x < len(image[0])+1; x++ {
      darkVoid := true
			if iterationNum == 0 {
				darkVoid = true
			} else if iterationNum%2 == 1 && algorithm[0] == '#' {
				darkVoid = false
			} else if iterationNum%2 == 0 && algorithm[0] == '#' {
				darkVoid = algorithm[len(algorithm) - 1] == '.'
			}
			pixelSurroundings := getPixelSurroundings(x, y, image, darkVoid)
			ret[y+1] = ret[y+1] + string(algorithm[pixels2Num(pixelSurroundings)])
		}
	}
	return ret
}

func countPixels(image []string, pixel rune) int {
	var ret int
	for _, line := range image {
		for _, c := range line {
			if c == '#' {
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
	// consume empty space
	scanner.Scan()
	inputImage := readImage(scanner)

	var outputImage []string = inputImage
  var part1, part2 int
	for i := 0; i < 50; i++ {
		outputImage = enhaceImage(outputImage, algorithm, i)
    if i == 1 {
      part1 = countPixels(outputImage, '#')
    } else if i == 49 {
      part2 = countPixels(outputImage, '#')
    }
	}
  fmt.Println(part1)
  fmt.Println(part2)
}
