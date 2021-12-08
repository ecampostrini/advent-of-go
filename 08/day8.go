package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Entry struct {
	SignalPatterns map[int][]string
	OutputValues   []string
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

func SortString(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func parseEntries(scanner *bufio.Scanner) []Entry {
	var ret []Entry
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		entry := Entry{make(map[int][]string), SplitAndTrim(parts[1], " ")}
		for _, pattern := range SplitAndTrim(parts[0], " ") {
			entry.SignalPatterns[len(pattern)] = append(entry.SignalPatterns[len(pattern)], SortString(pattern))
		}
		ret = append(ret, entry)
	}
	return ret
}

func permutations(input string) []string {
	var ret []string

	var helper func([]rune, int)
	helper = func(perm []rune, pos int) {
		if pos >= len(perm) {
			ret = append(ret, string(perm))
			return
		}

		for i := pos; i < len(perm); i++ {
			perm[pos], perm[i] = perm[i], perm[pos]
			helper(perm, pos+1)
			perm[pos], perm[i] = perm[i], perm[pos]
		}
	}

	helper([]rune(input), 0)

	return ret
}

var displayConfig = map[int][]int{
	0: {0, 1, 2, 4, 5, 6},
	1: {2, 5},
	2: {0, 2, 3, 4, 6},
	3: {0, 2, 3, 5, 6},
	4: {1, 2, 3, 5},
	5: {0, 1, 3, 5, 6},
	6: {0, 1, 3, 4, 5, 6},
	7: {0, 2, 5},
	8: {0, 1, 2, 3, 4, 5, 6},
	9: {0, 1, 2, 3, 5, 6},
}

func getConfigForDigit(digit int, config string) string {
	var ret string
	for _, pos := range displayConfig[digit] {
		ret += config[pos : pos+1]
	}
	return SortString(ret)
}

func main() {
	file, err := os.Open("./input.txt")
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

	configurations := permutations("abcdefg")
	var total int
	for _, entry := range entries {
		patterns := entry.SignalPatterns
		var targetConfig string
		for _, configuration := range configurations {
			if one := getConfigForDigit(1, configuration); one != patterns[2][0] {
				continue
			}
			if getConfigForDigit(7, configuration) != patterns[3][0] {
				continue
			}
			if four := getConfigForDigit(4, configuration); four != patterns[4][0] {
				continue
			}
			if eight := getConfigForDigit(8, configuration); eight != patterns[7][0] {
				continue
			}

			zero, six, nine := getConfigForDigit(0, configuration), getConfigForDigit(6, configuration), getConfigForDigit(9, configuration)
			hasMismatch := false
			for i := 0; !hasMismatch && i < len(patterns[6]); i++ {
				p := patterns[6][i]
				if p != zero && p != six && p != nine {
					hasMismatch = true
				}
			}
			if hasMismatch {
				continue
			}

			two, three, five := getConfigForDigit(2, configuration), getConfigForDigit(3, configuration), getConfigForDigit(5, configuration)
			hasMismatch = false
			for i := 0; !hasMismatch && i < len(patterns[5]); i++ {
				p := patterns[5][i]
				if p != two && p != three && p != five && p != nine {
					hasMismatch = true
				}
			}
			if hasMismatch {
				continue
			}

			targetConfig = configuration
			break
		}
		var outputValue int
		for _, outVal := range entry.OutputValues {
			for i := 0; i < 10; i++ {
				if getConfigForDigit(i, targetConfig) == SortString(outVal) {
					outputValue = outputValue*10 + i
					break
				}
			}
		}
		total += outputValue
	}
	fmt.Println(total)
}
