package slices

import (
	"fmt"
)

func PrintGridString(grid [][]string) {
	for _, row := range grid {
		for _, c := range row {
			if c == "" {
				fmt.Printf(".")
			} else {
				fmt.Printf("%s", c)
			}
		}
		fmt.Printf("\n")
	}
}

func MakeBidimensionalSliceString(dx, dy int) [][]string {
	ret := make([][]string, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]string, dx)
	}
	return ret
}
