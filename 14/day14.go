package main

import (
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

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	scanner.Scan()
	template := scanner.Text()
	//fmt.Println("Template: ", template)

	scanner.Scan()
	insertionRules := ParseStringMap(scanner, "->")
	//fmt.Println(insertionRules)

	occurrenceMap := make(map[string]int)
	for _, c := range template {
		occurrenceMap[string(c)]++
	}

	for step := 0; step < 40; step++ {
    fmt.Printf("Step: %d, template size: %d\n", step, len(template))
		currentPolymer := template[0:1]
		for i := 0; i < len(template)-1; i++ {
			currentPair := template[i : i+2]
			newElement := insertionRules[currentPair]
			currentPolymer = currentPolymer + newElement + currentPair[1:2]
			occurrenceMap[newElement]++
		}

		template = currentPolymer
		//fmt.Printf("Occurrence map: %v\n", occurrenceMap)
		//fmt.Printf("Step %d: %s\nLength: %d", step, template, len(template))
	}

	max, min := -1, int(1e10)
	for _, count := range occurrenceMap {
		if count < min {
      min = count
		}

    if count > max {
      max = count
    }
	}

  fmt.Printf("\nPart1: %d\n", max - min)
}
