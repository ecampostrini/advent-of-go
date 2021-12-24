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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	xBounds, yBounds := parseBounds(scanner)

	var yCandidates = make(map[int]bool)
	for i := 0; i < 1000; i++ {
		y, velocity := i, i
		var steps []int
		for j := 0; j < 1000; j++ {
			steps = append(steps, y)
			if y >= yBounds.A && y <= yBounds.B && !yCandidates[i] {
				yCandidates[i] = true
				fmt.Printf("iv: %v, steps: %v\n", i, steps)
			}
			velocity--
			y += velocity
		}
	}
	fmt.Printf("Y-candidates: %v\n", yCandidates)
	fmt.Println("---")
	var xCandidates = make(map[int]bool)
	for i := 0; i < 1000; i++ {
		x, velocity := i, i
		var steps []int
		for j := 0; j < 1000; j++ {
			steps = append(steps, x)
			if x >= xBounds.A && x <= xBounds.B && !xCandidates[i] {
				xCandidates[i] = true
				fmt.Printf("iv: %v, steps: %v\n", i, steps)
			}
			if velocity > 0 {
				velocity--
			}
			x += velocity
		}
	}
	fmt.Printf("x-candidates: %v\n", xCandidates)

	var candidates = make(map[Point][]Point)
	for ivX, _ := range xCandidates {
		for ivY, _ := range yCandidates {
			var steps []Point
			candidate := Point{ivX, ivY}
			x, xVel := ivX, ivX
			y, yVel := ivY, ivY
			for step := 0; step < 10000; step++ {
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
				x += xVel
				yVel--
				y += yVel

			}
		}
	}

	maxY := math.MinInt
	var bestInitialVel Point
	for initialVel, steps := range candidates {
		for _, p := range steps {
			if p.Y > maxY {
				maxY = p.Y
				bestInitialVel = initialVel
			}
		}
	}
	fmt.Printf("%v: %d\n", bestInitialVel, maxY)
}
