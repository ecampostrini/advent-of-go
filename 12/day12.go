package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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

func isLowerCase(s string) bool {
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func countPaths(graph Graph) int {
	var helper func(string)
	var pathCount int
	currentPath := []string{"start"}
	visited := make(map[string]bool)

	helper = func(currentNode string) {
		if currentNode == "end" {
			pathCount++
			fmt.Printf("%v\n", currentPath)
			return
		}

		if isLowerCase(currentNode) {
			visited[currentNode] = true
		}
		for _, neighbour := range graph[currentNode] {
			if visited[neighbour] {
				continue
			}
			currentPath = append(currentPath, neighbour)
			helper(neighbour)
			currentPath = currentPath[:len(currentPath)-1]
		}
		visited[currentNode] = false
	}

	helper("start")

	return pathCount
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to read input file: ", err)
		os.Exit(1)
	}
	graph := parseGraph(bufio.NewScanner(file))
	fmt.Printf("%v\n", graph)

	fmt.Println(countPaths(graph))
}
