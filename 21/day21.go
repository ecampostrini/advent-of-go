package main

import (
	"bufio"
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	//"github.com/ecampostrini/advent-of-go/utils/slices"
	"strconv"
	"strings"
)

func getDeterministicDice() func() int {
	value := 1
	ret := func() int {
		ret := value
		value = (value % 100) + 1
		return ret
	}
	return ret
}

func parsePositions(scanner *bufio.Scanner) []int {
	var ret []int
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		position, err := strconv.Atoi(fields[len(fields)-1])
		if err != nil {
			panic(fmt.Sprintf("Failed to parse initial positions: %v", err))
		}
		ret = append(ret, position)
	}
	return ret
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	positions := parsePositions(scanner)
  dice := getDeterministicDice()
  score := make([]int, len(positions))
  var winner, diceRollCount int
  hasWinner := false
  for !hasWinner {
    for turn := 0; !hasWinner && turn < len(positions); turn++ {
      diceSum := dice() + dice() + dice()
      newPosition := (positions[turn] + diceSum) % 10
      if newPosition == 0 {
        newPosition = 10
      }
      positions[turn] = newPosition
      score[turn] += positions[turn]
      diceRollCount += 3
      if score[turn] >= 1000 {
        hasWinner = true
        winner = turn
      }
    }
  }
  loser := (winner + 1) % 2
  fmt.Println(diceRollCount * score[loser])
}
