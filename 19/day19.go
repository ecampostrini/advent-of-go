package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"strconv"
	"strings"
)

type Point2d struct {
	X, Y int
}

type Point3d struct {
	X, Y, Z int
}

type ScannerReport2d = map[Point2d]bool
type ScannerSlice2d = []Point2d
type ScannerReport = map[Point3d]bool

func rotate2d(p Point2d, amount int) Point2d {
	if amount == 0 {
		return p
	} else {
		p.X, p.Y = p.Y, p.X
		return p
	}
}

func parseInput2d(scanner *bufio.Scanner) map[string][]Point2d {
	ret := make(map[string][]Point2d)
	var currentScanner string
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if strings.Contains(line, "scanner") {
			currentScanner = strings.TrimSuffix(strings.TrimPrefix(line, "--- "), " ---")
			scanner.Scan()
			line = scanner.Text()
		}

		coords := strings.Split(line, ",")
		var x, y int
		var err error
		if x, err = strconv.Atoi(coords[0]); err != nil {
			panic(fmt.Sprintf("Failed to parse number: %v", err))
		}
		if y, err = strconv.Atoi(coords[1]); err != nil {
			panic(fmt.Sprintf("Failed to parse number: %v", err))
		}
		ret[currentScanner] = append(ret[currentScanner], Point2d{x, y})
	}
	return ret
}

func getOverlappingPoints2d(r1, r2 []Point2d) (int, Point2d, bool) {
	r2Map := make(map[Point2d]bool)
	for _, p := range r2 {
		r2Map[p] = true
	}

	for x := -1; x < 2; x += 2 {
		for y := -1; y < 2; y += 2 {
			for rotation := 0; rotation < 2; rotation++ {
				r1p0 := r1[0]
				for _, r2p := range r2 {
					r2p = rotate2d(r2p, rotation)
					r2p.X *= x
					r2p.Y *= y
					candidate := Point2d{r1p0.X - r2p.X, r1p0.Y - r2p.Y}
					var matchCount int = 0
					for i := 0; matchCount < 3 && i < len(r1); i++ {
						transformed := Point2d{r1[i].X - candidate.X, r1[i].Y - candidate.Y}
						if r2Map[transformed] {
							matchCount++
						}
					}
					if matchCount == 3 {
						return matchCount, candidate, true
					}
				}

			}
		}
	}
	return 0, Point2d{0, 0}, false
}

func main() {
	scanner, file := files.ReadFile("./input2d.test.txt")
	defer file.Close()

	readings := parseInput2d(scanner)

	c, p, ok := getOverlappingPoints2d(readings["scanner 0"], readings["scanner 1"])
	if ok {
		fmt.Printf("%d: %v\n", c, p)
	}

	fmt.Printf("%v: %d\n", readings, len(readings))
}
