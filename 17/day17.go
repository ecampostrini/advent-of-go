package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"math"
	"strconv"
	"strings"
)

type Bounds struct {
	A, B int
}

type Point struct {
	X, Y int
}

func parseBounds(scanner *bufio.Scanner) (Bounds, Bounds) {
	scanner.Scan()
	input := strings.Replace(scanner.Text(), "target area: ", "", 1)
	input = strings.Replace(input, "x=", "", 1)
	input = strings.Replace(input, "y=", "", 1)
	bounds := strings.Split(input, ",")

	xBounds := strings.Split(strings.TrimSpace(bounds[0]), "..")
	x1, err1 := strconv.Atoi(xBounds[0])
	x2, err2 := strconv.Atoi(xBounds[1])
	if err1 != nil || err2 != nil {
		panic(fmt.Sprintf("Error while parsing bounds: %v, %v", err1, err2))
	}
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	yBounds := strings.Split(strings.TrimSpace(bounds[1]), "..")
	y1, err1 := strconv.Atoi(yBounds[0])
	y2, err2 := strconv.Atoi(yBounds[1])
	if err1 != nil || err2 != nil {
		panic(fmt.Sprintf("Error while parsing bounds: %v, %v", err1, err2))
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	return Bounds{x1, x2}, Bounds{y1, y2}
}

func isWithinBounds(n int, bounds Bounds) bool {
	return n >= bounds.A && n <= bounds.B
}

func findCandidates(xBounds, yBounds Bounds) map[Point][]Point {
	var candidates = make(map[Point][]Point)
	for ivX := 0; ivX < 1000; ivX++ {
		for ivY := -1000; ivY < 1000; ivY++ {
			candidate := Point{ivX, ivY}
			var steps []Point
			x, xVel := 0, ivX
			y, yVel := 0, ivY
			for step := 0; step < 10000; step++ {
				x, y = x+xVel, y+yVel
				steps = append(steps, Point{x, y})
				if isWithinBounds(x, xBounds) && isWithinBounds(y, yBounds) {
					candidates[candidate] = steps
					break
				} else if x > xBounds.B || y < yBounds.A {
					break
				}
				if xVel > 0 {
					xVel--
				}
				yVel--
			}
		}
	}
	return candidates
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	xBounds, yBounds := parseBounds(scanner)
	candidates := findCandidates(xBounds, yBounds)
	maxY := math.MinInt
  //var bestInitialVel Point
	for _, steps := range candidates {
		for _, step := range steps {
			if step.Y > maxY {
				maxY = step.Y
				//bestInitialVel = initialVel
			}
		}
	}
	// for debugging purposes
	//fmt.Printf("%v\n", bestInitialVel)
	// part1
	fmt.Printf("%v\n", maxY)
	// part2
	fmt.Printf("%d\n", len(candidates))
}
