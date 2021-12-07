package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) uint {
	if x < 0 {
		return uint(-x)
	}
	return uint(x)
}

func Max(s []int) int {
	max := 0
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

func sumFirstN(n uint) uint {
	return n * (n + 1) / 2
}

func readPositions(scanner *bufio.Scanner) []int {
	var ret []int
	scanner.Scan()
	for _, c := range strings.Split(scanner.Text(), ",") {
		pos, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Error while parsing position: ", err)
			os.Exit(1)
		}
		ret = append(ret, pos)
	}
	return ret
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	minFuel := ^uint(0) >> 1
	positions := readPositions(bufio.NewScanner(file))
	max := Max(positions)
	for i := 0; i <= max; i++ {
		currentFuel := uint(0)
		for _, pos := range positions {
			currentFuel += sumFirstN(Abs(i - pos))
		}
		if currentFuel < minFuel {
			minFuel = currentFuel
		}
	}

	fmt.Println(minFuel)
}
