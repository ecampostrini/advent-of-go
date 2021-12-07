package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func initLanternFish(s *bufio.Scanner, cycleLength int) map[int]int {
	ret := make(map[int]int)
	for i := 0; i < cycleLength+1; i++ {
		ret[i] = 0
	}

	s.Scan()
	for _, c := range strings.Split(s.Text(), ",") {
		n, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Failed to parse number from input", err)
			os.Exit(1)
		}
		ret[n]++
	}
	return ret
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening input file", err)
		os.Exit(1)
	}
	defer file.Close()

	const days = 256
	lanternfish := initLanternFish(bufio.NewScanner(file), 8)
	for dayNum := 0; dayNum < days; dayNum++ {
		timeouts := lanternfish[0]
		for j := 0; j < len(lanternfish)-1; j++ {
			lanternfish[j] = lanternfish[j+1]
		}
		lanternfish[len(lanternfish)-1] = timeouts
		lanternfish[6] += timeouts
	}

	count := 0
	for _, v := range lanternfish {
		count += v
	}

	fmt.Println(count)
}
