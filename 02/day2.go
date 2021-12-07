package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMovements(filePath string) ([]string, error) {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func main() {
	movements, _ := readMovements("./input.txt")

	//part 1
	var x, y = 0, 0
	for i := 0; i < len(movements); i++ {
		m := strings.Split(movements[i], " ")
		direction := m[0]
		amount, _ := strconv.Atoi(m[1])
		if direction == "forward" {
			x += amount
		} else if direction == "up" {
			y -= amount
		} else {
			y += amount
		}
	}

	fmt.Println(x, y, x*y)

	// part 2
	var hor, aim, depth = 0, 0, 0
	for _, mov := range movements {
		m := strings.Split(mov, " ")
		direction := m[0]
		amount, _ := strconv.Atoi(m[1])
		if direction == "forward" {
			hor += amount
			depth += aim * amount
		} else if direction == "up" {
			aim -= amount
		} else {
			aim += amount
		}
	}
	fmt.Println(hor, depth, hor*depth)
}
