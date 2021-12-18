package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"math"
)

type Point struct {
	X, Y int
}

type WeightedPoint struct {
	p Point
	w int
}

// heap implementation taken from https://pkg.go.dev/container/heap
type MinHeap []WeightedPoint

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].w <= h[j].w }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(WeightedPoint))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}

func ParseGridInt(scanner *bufio.Scanner) [][]int {
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func doit(grid [][]int) int {
	grid[0][0] = 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if i == 0 && j == 0 {
				continue
			}
			upVal, leftVal := math.MaxInt, math.MaxInt
			if i > 0 {
				upVal = grid[i-1][j]
			}
			if j > 0 {
				leftVal = grid[i][j-1]
			}
			grid[i][j] += Min(upVal, leftVal)
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func neighbours(p Point, grid [][]int) []Point {
	var ret []Point
	px, py := p.X, p.Y
	if px > 0 {
		ret = append(ret, Point{px - 1, py})
	}
	if px < len(grid[0])-1 {
		ret = append(ret, Point{px + 1, py})
	}
	if py > 0 {
		ret = append(ret, Point{px, py - 1})
	}
	if py < len(grid)-1 {
		ret = append(ret, Point{px, py + 1})
	}
	return ret
}

func expandGrid(grid [][]int, dimension int) [][]int {
	dy, dx := len(grid), len(grid[0])
	newGrid := make([][]int, dy*dimension)
	for i, _ := range newGrid {
		newGrid[i] = make([]int, dx*dimension)
	}

	for i, _ := range newGrid {
		for j, _ := range newGrid[i] {
			oi, oj := i%dy, j%dx
			offsetY, offsetX := i/dy, j/dx
			value := grid[oi][oj] + offsetY + offsetX
			newGrid[i][j] = (value / 10) + (value % 10)
		}
	}
	return newGrid
}

func calculateShortestPath(grid [][]int) int {
	distanceMap := map[Point]int{Point{0, 0}: 0}
	minHeap := &MinHeap{WeightedPoint{Point{0, 0}, 0}}
	heap.Init(minHeap)
	for minHeap.Len() > 0 {
		currentPoint := minHeap.Pop().(WeightedPoint)
		for _, n := range neighbours(currentPoint.p, grid) {
			nDistance, ok := distanceMap[n]
			if !ok {
				nDistance = math.MaxInt
			}
			newDistance := distanceMap[currentPoint.p] + grid[n.Y][n.X]
			if newDistance < nDistance {
				distanceMap[n] = newDistance
				heap.Push(minHeap, WeightedPoint{n, newDistance})
			}
		}
	}
	return distanceMap[Point{len(grid[0]) - 1, len(grid) - 1}]
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	grid := ParseGridInt(scanner)
	// part 1
	fmt.Println(calculateShortestPath(grid))
	// part 2
	fmt.Println(calculateShortestPath(expandGrid(grid, 5)))
}
