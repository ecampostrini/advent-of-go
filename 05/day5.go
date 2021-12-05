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

func readSegments(scanner *bufio.Scanner) []Segment {
  var ret []Segment
  for scanner.Scan() {
    rawPoints := strings.Split(scanner.Text(), "->") 
    ret = append(ret, Segment{parsePoint(rawPoints[0]), parsePoint(rawPoints[1])})
  }
  return ret
}
// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDistanceAndDirection(segment *Segment) (int, Direction) {
  var distance int
  var direction Direction

  if segment.a.X == segment.b.X {
    difference := segment.a.Y - segment.b.Y
    if difference >= 0 {
      direction = Direction{0, -1}
    } else {
      direction = Direction{0, 1}
    }
    distance = Abs(difference)
  } else if segment.a.Y == segment.b.Y {
    difference := segment.a.X - segment.b.X
    if difference >= 0 {
      direction = Direction{-1, 0}
    } else {
      direction = Direction{1, 0}
    }
    distance = Abs(difference)
  }

  return distance, direction
}

func main() {
  file, err := os.Open("./input.test.txt")
  if err != nil {
    fmt.Println("Oops")
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  segments := readSegments(scanner)

  for i := 0; i < len(segments); i++ {
    fmt.Println("---")
    fmt.Println(segments[i])
    fmt.Println(getDistanceAndDirection(&segments[i]))
  }
}
