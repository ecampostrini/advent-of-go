package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph = map[string][]string

var graph = make(Graph)

func parseGraph(scanner *bufio.Scanner) Graph {
	ret := make(Graph)
	for scanner.Scan() {
		vertices := strings.Split(scanner.Text(), "-")
		ret[vertices[0]] = append(ret[vertices[0]], vertices[1])
		if vertices[0] != "start" && vertices[1] != "end" {
			ret[vertices[1]] = append(ret[vertices[1]], vertices[0])
		}
	}
	return ret
}

func main() {
	file, err := os.Open("./input.test.txt")
	if err != nil {
		fmt.Println("Failed to read input file: ", err)
		os.Exit(1)
	}
	graph := parseGraph(bufio.NewScanner(file))

	fmt.Printf("%v\n", graph)
}
