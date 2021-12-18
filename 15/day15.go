package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"github.com/ecampostrini/advent-of-go/utils/slices"
	"math"
)

type Point struct {
	X, Y int
}

type WeightedPoint struct {
	p Point
	w int
}

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
  //n := len(old)
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

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	grid := ParseGridInt(scanner)
	//slices.PrintGridInt(grid)

	const dimension = 5
	dy, dx := len(grid), len(grid[0])
	fmt.Printf("dy: %d, dx: %d\n", dy*dimension, dx*dimension)
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

	fmt.Println("\n---\n")
	slices.PrintGridInt(newGrid, "")

	visited := make(map[Point]bool)
	distanceMap := map[Point]int{Point{0, 0}: 0}
	minHeap := &MinHeap{WeightedPoint{Point{0, 0}, 0}}
	heap.Init(minHeap)
	for minHeap.Len() > 0 {
		currentPoint := minHeap.Pop().(WeightedPoint)
		//fmt.Printf("Current point: %v\n", currentPoint)
		if visited[currentPoint.p] {
      //continue
		}
		//if currentPoint.p.X == len(newGrid[0])-1 && currentPoint.p.Y == len(newGrid)-1 {
		//break
		//}
		for _, n := range neighbours(currentPoint.p, newGrid) {
			//fmt.Printf("  neighbour:%v\n", n)
			nDistance, ok := distanceMap[n]
			if !ok {
				nDistance = math.MaxInt
			}
			newDistance := distanceMap[currentPoint.p] + newGrid[n.Y][n.X]
			if newDistance < nDistance {
				distanceMap[n] = newDistance
				heap.Push(minHeap, WeightedPoint{n, newDistance})
			}
		}
		visited[currentPoint.p] = true
		//fmt.Printf("    distancemap: %v\n", distanceMap)
		//fmt.Printf("    heap: %v\n", minHeap)
	}

	fmt.Println(distanceMap[Point{len(newGrid[0]) - 1, len(newGrid) - 1}])
	//fmt.Println("\n---\n")
	//slices.PrintGridInt(newGrid, " ")

}
