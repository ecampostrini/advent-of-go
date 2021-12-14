package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"math"
	"strings"
)

func ParseStringMap(scanner *bufio.Scanner, separator string) map[string]string {
	ret := make(map[string]string)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), separator)
		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		ret[key] = value
	}
	return ret
}

func SumValuesInt(m map[string]int) int {
	var ret int
	for _, count := range m {
		ret += count
	}
	return ret
}

func doit(template string, insertionRules map[string]string, iterations int) int {
	polymerMap := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		polymerMap[template[i:i+2]]++
	}

	occurrenceMap := make(map[string]int)
	for _, c := range template {
		occurrenceMap[string(c)]++
	}

	for i := 0; i < iterations; i++ {
		newPolymerMap := make(map[string]int)
		for polymer, count := range polymerMap {
			newElement := insertionRules[polymer]
			newPolymerMap[string(polymer[0])+newElement] += count
			newPolymerMap[newElement+string(polymer[1])] += count
			occurrenceMap[newElement] += count
		}
		polymerMap = newPolymerMap
	}

	max, min := -1, math.MaxInt
	for _, count := range occurrenceMap {
		if count < min {
			min = count
		}

		if count > max {
			max = count
		}
	}
	return max - min
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()
	scanner.Scan()
	template := scanner.Text()
	scanner.Scan()
	insertionRules := ParseStringMap(scanner, "->")
	// part 1
	fmt.Println(doit(template, insertionRules, 10))
	// part 2
	fmt.Println(doit(template, insertionRules, 40))
}
