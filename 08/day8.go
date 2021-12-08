package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	SignalPatterns, OutputValues []string
}

func SplitAndTrim(input, delim string) []string {
	var ret []string
	for _, tok := range strings.Split(input, delim) {
		if trimmedTok := strings.TrimSpace(tok); len(trimmedTok) > 0 {
			ret = append(ret, trimmedTok)
		}
	}
	return ret
}

func parseEntries(scanner *bufio.Scanner) []Entry {
	var ret []Entry
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		ret = append(ret, Entry{SplitAndTrim(parts[0], " "), SplitAndTrim(parts[1], " ")})
	}
	return ret
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening input file: ", err)
	}
	defer file.Close()

	entries := parseEntries(bufio.NewScanner(file))
	easyDigitsCount := 0
	for _, entry := range entries {
		for _, digit := range entry.OutputValues {
			digitLen := len(digit)
			if digitLen == 2 || digitLen == 3 || digitLen == 4 || digitLen == 7 {
				easyDigitsCount++
			}
		}
	}
	fmt.Println(easyDigitsCount)
}
