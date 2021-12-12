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

func countPaths(graph Graph) int {
	var helper func(string)
	var pathCount int
	currentPath := []string{"start"}
	maxVisitCount := make(map[string]int)
	visitedPaths := make(map[string]bool)

	helper = func(currentNode string) {
		if currentNode == "end" {
			pathAsKey := strings.Join(currentPath, "")
			if !visitedPaths[pathAsKey] {
				pathCount++
				visitedPaths[pathAsKey] = true
			}
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

	for _, cave := range getLowerCaseCaves(graph) {
		maxVisitCount[cave]++
		helper("start")
		maxVisitCount[cave]--
	}

	return pathCount
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to read input file: ", err)
		os.Exit(1)
	}
	graph := parseGraph(bufio.NewScanner(file))
	fmt.Println(countPaths(graph))
}
