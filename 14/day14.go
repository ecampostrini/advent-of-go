package main

import (
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"github.com/ecampostrini/advent-of-go/utils/slices"
)

func main() {
	scanner, file := files.ReadFile("./input.test.txt")
	defer file.Close()

	scanner.Scan()
	template := scanner.Text()
	fmt.Println("Template: ", template)
	insertionRules := slices.ParseStringTuple(scanner, "->")
	slices.PrintSliceStringTuple(insertionRules, "->")

}
