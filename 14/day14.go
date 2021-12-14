package main

import (
  "math"
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
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

func getLength(m map[string]int) int {
  var ret int
  for _, count := range m {
    ret += count
  }
  return ret
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	scanner.Scan()
	template := scanner.Text()
	//fmt.Println("Template: ", template)

	scanner.Scan()
	insertionRules := ParseStringMap(scanner, "->")
	//fmt.Println(insertionRules)

	polymerMap := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		polymerMap[template[i:i+2]]++
	}
	fmt.Printf("Polymer map: %v\n", polymerMap)

	occurrenceMap := make(map[string]int)
	for _, c := range template {
		occurrenceMap[string(c)]++
	}
	fmt.Printf("Occurrence map: %v\n", occurrenceMap)

	for i := 0; i < 40; i++ {
		newPolymerMap := make(map[string]int)
		for polymer, count := range polymerMap {
			newElement := insertionRules[polymer]
			newPolymerMap[string(polymer[0])+newElement] += count
			newPolymerMap[newElement+string(polymer[1])] += count
			occurrenceMap[newElement] += count
		}
		polymerMap = newPolymerMap
    fmt.Printf("%d - Polymer map: %v\n", i, polymerMap)
    fmt.Printf("Occurrence map: %v\n", occurrenceMap)
    fmt.Printf("polymenr length: %d\n", getLength(occurrenceMap))
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

	fmt.Printf("\nPart1: %d\n", max-min)
}
