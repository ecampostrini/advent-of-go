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
type ScannerReport3d = map[Point3d]bool
type ScannerSlice3d = []Point3d

func rotate2d(p Point2d, amount int) Point2d {
	if amount == 1 {
		p.X, p.Y = p.Y, p.X
		return p
	}
	return p
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

func parseInput3d(scanner *bufio.Scanner) map[string][]Point3d {
	ret := make(map[string][]Point3d)
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
		var x, y, z int
		var err error
		if x, err = strconv.Atoi(coords[0]); err != nil {
			panic(fmt.Sprintf("Failed to parse number: %v", err))
		}
		if y, err = strconv.Atoi(coords[1]); err != nil {
			panic(fmt.Sprintf("Failed to parse number: %v", err))
		}
		if z, err = strconv.Atoi(coords[2]); err != nil {
			panic(fmt.Sprintf("Failed to parse number: %v", err))
		}
		ret[currentScanner] = append(ret[currentScanner], Point3d{x, y, z})
	}
	return ret
}

func rotate3d(p Point3d, amount int) Point3d {
	if amount == 1 {
		p.X, p.Y, p.Z = p.Y, p.Z, p.X
	} else if amount == 2 {
		p.X, p.Y, p.Z = p.Z, p.X, p.Y
	} else if amount == 3 {
		p.X, p.Y, p.Z = p.X, p.Z, p.Y
	} else if amount == 4 {
		p.X, p.Y, p.Z = p.Y, p.X, p.Z
	} else if amount == 5 {
		p.X, p.Y, p.Z = p.Z, p.Y, p.X
	}
	return p
}

func getOverlappingPoints3d(r1, r2 []Point3d) (int, Point3d, bool) {
	r2Map := make(map[Point3d]bool)
	for _, p := range r2 {
		r2Map[p] = true
	}
	var maxMatchCount int = -1
	var bestCandidate Point3d
	var foundCandidate bool
	for x := -1; x < 2; x += 2 {
		for y := -1; y < 2; y += 2 {
			for z := -1; z < 2; z += 2 {
				for rotation := 0; rotation < 6; rotation++ {
					//r1p0 := r1[0]
					for _, r1p := range r1 {
						for _, r2p := range r2 {
							r2p = rotate3d(r2p, rotation)
							r2p.X *= x
							r2p.Y *= y
							r2p.Z *= z
              candidate := Point3d{r1p.X + r2p.X, r1p.Y + r2p.Y, r1p.Z + r2p.Z}
							//if candidate.X == 68 {
							//fmt.Printf("Candidate: %v\n", candidate)
							//}
							var matchCount int = 0
							for i := 0; i < len(r1); i++ {
                r1t := r1[i]
                //r1t := rotate3d(r1[i], rotation)
                //r1t.X *= x
                //r1t.Y *= y
                //r1t.Z *= z
                transformed := Point3d{(-x) * (r1t.X - candidate.X), (-y) * (r1t.Y - candidate.Y), (-z) * (r1t.Z - candidate.Z)}
                transformed = rotate3d(transformed, rotation)
                //transformed := Point3d{(r1t.X + candidate.X), (r1t.Y + candidate.Y),(r1t.Z + candidate.Z)}
								if r2Map[transformed] {
									matchCount++
                  //if candidate.X == 68 {
                    //fmt.Printf("Here: %v\n", r2[i])
                  //}
								}
							}
							if matchCount > maxMatchCount {
								maxMatchCount = matchCount
								bestCandidate = candidate
								foundCandidate = true
							}
						}
					}
				}
			}
		}
	}
	return maxMatchCount, bestCandidate, foundCandidate
}

func main() {
	scanner, file := files.ReadFile("./input3d.test.txt")
	defer file.Close()

	readings := parseInput3d(scanner)

	c, p, ok := getOverlappingPoints3d(readings["scanner 0"], readings["scanner 4"])
	if ok {
		fmt.Printf("%d: %v\n", c, p)
	}

	var transformedSlice []Point3d
	for _, point := range readings["scanner 1"] {
		transformedPoint := Point3d{(point.X - p.X), (point.Y - p.Y), (point.Z - p.Z)}
		transformedSlice = append(transformedSlice, transformedPoint)
	}
	fmt.Printf("%v\n", transformedSlice)
	//readings["scanner 1"] = transformedSlice
	//c, p, ok = getOverlappingPoints3d(readings["scanner 1"], readings["scanner 4"])
	//if ok {
	//fmt.Printf("%d: %v\n", c, p)
	//}

	//fmt.Printf("%v: %d\n", readings, len(readings))
}
