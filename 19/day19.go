package main

import (
	"fmt"
)

type Point2d struct {
	X, Y int
}

type Point3d struct {
	X, Y, Z int
}

type ScannerReport = map[Point3d]bool

func rotate2d(p Point2d, int amount) Point2d {
  if amount == 0 {
    return p
  } else 
    p.X, p.Y = p.Y, p.X
    return p
  }
}

func getOverlappingPoints2d(r1, r2 ScannerReport) ([]Point2d, Point2d) {
	for x := -1; x < 2; x+=2 {
		for y := -1; y < 2; y+=2 {
			for rotation := 0; rotation < 2; rotation++ {
        for id1, r2p := range r2 {
          r2p = rotate2d(r2p, rotation)
          r2p.X *= x 
          r2p.Y *= y
          for id2, r1p := range r1 {
            candidate := Point2d{r1p.X - r2p.X, r1p.Y - r2p.Y}
          }
        }
			}
		}
	}
}

//func getOverlappingPoints3d(report1, report2 ScannerReport) ([]Point3d, Point3d) {
//for x := 0; x < 2; x++ {
//for y := 0; y < 2; y++ {
//for z := 0; z < 2; z++ {
//for rotation := 0; rotation < 3; rotation++ {

//}
//}
//}
//}
//}

func main() {

}
