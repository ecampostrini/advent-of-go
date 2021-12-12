package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getLowerCaseCaves(g Graph) []string {
	var ret []string

	for k, _ := range g {
		if isLowerCase(k) && k != "start" && k != "end" {
			ret = append(ret, k)
		}
	}
	sort.Strings(ret)
	return ret
}

func getPaths(graph Graph, maxVisitCount map[string]int) []string {
	var helper func(string)
	var paths []string
	currentPath := []string{"start"}

	helper = func(currentNode string) {
		if currentNode == "end" {
			paths = append(paths, strings.Join(currentPath, ""))
			return
		}

		if isLowerCase(currentNode) {
			maxVisitCount[currentNode] -= 1
		}

		for _, neighbour := range graph[currentNode] {
			if isLowerCase(neighbour) && maxVisitCount[neighbour] < 0 {
				continue
			}
			currentPath = append(currentPath, neighbour)
			helper(neighbour)
			currentPath = currentPath[:len(currentPath)-1]
		}
		maxVisitCount[currentNode] += 1
	}

	helper("start")

	return paths
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to read input file: ", err)
		os.Exit(1)
	}

	graph := parseGraph(bufio.NewScanner(file))

	maxVisitCount := make(map[string]int)
	// part 1
	fmt.Println(len(getPaths(graph, maxVisitCount)))

	// part 2
	visitedPaths := make(map[string]bool)
	for _, cave := range getLowerCaseCaves(graph) {
		maxVisitCount[cave]++
		for _, path := range getPaths(graph, maxVisitCount) {
			if !visitedPaths[path] {
				visitedPaths[path] = true
			}
		}
		maxVisitCount[cave]--
	}
	fmt.Println(len(visitedPaths))
}
