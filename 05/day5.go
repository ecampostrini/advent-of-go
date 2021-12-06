package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

type Point struct {
  X int
  Y int
}

type Direction = Point

type Segment struct {
  a Point
  b Point
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func makeGrid(dim int) [][]int {
  ret := make([][]int, dim)
  for i := range ret {
    ret[i] = make([]int, dim)
  }
  return ret
}

func parsePoint(s string) Point {
  var x, y int
  var err error

  xy := strings.Split(s, ",")
  if x, err = strconv.Atoi(strings.TrimSpace(xy[0])); err != nil {
    fmt.Println("Failed to parse coordinate", err)
    os.Exit(1)
  }

  if y, err = strconv.Atoi(strings.TrimSpace(xy[1])); err != nil {
    fmt.Println("Failed to parse coordinate", err)
    os.Exit(1)
  }
  
  return Point{x, y}
}

func readSegments(scanner *bufio.Scanner) ([]Segment, int, int) {
  var segments []Segment
  var maxX, maxY int
  for scanner.Scan() {
    rawPoints := strings.Split(scanner.Text(), "->") 
    newSegment := Segment{parsePoint(rawPoints[0]), parsePoint(rawPoints[1])}
    maxX = max(max(newSegment.a.X, newSegment.b.X), maxX)
    maxY = max(max(newSegment.a.Y, newSegment.b.Y), maxY)
    segments = append(segments, newSegment)
  }
  return segments, maxX + 1, maxY + 1
}

func getDirectionAndLength(segment *Segment) (Direction, int) {
  var difference int
  var direction Direction

  if segment.a.X - segment.b.X == segment.a.Y - segment.b.Y {
    difference = segment.a.X - segment.b.X
    if segment.a.X < segment.b.X {
      direction = Direction{1, 1}
    } else {
      direction = Direction{-1, -1}
    }
  } else if segment.a.X - segment.b.X == segment.b.Y - segment.a.Y {
    difference = segment.a.X - segment.b.X
    direction = Direction{-1, 1}
    if segment.a.X < segment.b.X {
      direction = Direction{1, -1}
    } else {
      direction = Direction{-1, 1}
    }
  } else if segment.a.X == segment.b.X {
    difference = segment.a.Y - segment.b.Y
    if difference >= 0 {
      direction = Direction{0, -1}
    } else {
      direction = Direction{0, 1}
    }
  } else if segment.a.Y == segment.b.Y {
    difference = segment.a.X - segment.b.X
    if difference >= 0 {
      direction = Direction{-1, 0}
    } else {
      direction = Direction{1, 0}
    }
  }

  return direction, Abs(difference)
}

func printGrid(g *[][]int) {
  grid := *g
  for i := range grid {
    fmt.Println(grid[i])
  }
}

func countOverlaps(grid *[][]int) int {
  var ret int
  for _, line := range *grid {
    for _, point := range line {
      if point >= 2 { ret++ }
    }
  }

  return ret
}

func markSegmentOnGrid(grid *[][]int, start Point, direction Direction, length int) {
  if direction.X == 0 && direction.Y == 0 {
    return
  }

  for i := 0; i <= length; i++ {
    (*grid)[start.Y + (direction.Y * i)][start.X + (direction.X * i)]++
  }
}

func main() {
  file, err := os.Open("./input.txt")
  if err != nil {
    fmt.Println("Oops")
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  segments, maxX, maxY := readSegments(scanner)

  grid := makeGrid(max(maxX, maxY))
  for i := 0; i < len(segments); i++ {
    direction, segmentLength := getDirectionAndLength(&segments[i])
    markSegmentOnGrid(&grid, segments[i].a, direction, segmentLength)
  }
  //printGrid(&grid)
  fmt.Println(countOverlaps(&grid))

}
