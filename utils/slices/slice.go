package slices

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/types"
	"strings"
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

func PrintGridInt(grid [][]int, separator string) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%d%s", c, separator)
		}
		fmt.Printf("\n")
	}
}

func PrintSliceStringTuple(st []types.StringTuple, separator string) {
	fmt.Println("[")
	for _, tuple := range st {
		fmt.Printf("  %s %s %s,\n", tuple.First, separator, tuple.Second)
	}
	fmt.Println("]")
}

func MakeBidimensionalSliceString(dx, dy int) [][]string {
	ret := make([][]string, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]string, dx)
	}
	return ret
}

func ParseStringTuple(scanner *bufio.Scanner, separator string) []types.StringTuple {
	var ret []types.StringTuple
	scanner.Scan()
	for scanner.Scan() {
		rule := strings.Split(scanner.Text(), separator)
		ret = append(ret, types.StringTuple{strings.TrimSpace(rule[0]), strings.TrimSpace(rule[1])})
	}
	return ret
}
