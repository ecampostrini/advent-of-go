package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Board = [][]string

func readBoard(scanner *bufio.Scanner) Board {
  var ret Board
  for scanner.Text() == "" {
    scanner.Scan()
  }
  for scanner.Text() != "" {
    ret = append(ret, strings.Split(scanner.Text(), " "))
    scanner.Scan()
  }
  return ret
}

func markNumber(board *Board, number string) {

}

func hasFullCol(board *Board) bool {
  return false
}

func hasFullRow(board *Board) bool {
  return false
}

func main() {
  file, err := os.Open("./input.txt")
  if err != nil {
    fmt.Println("Oops")
    os.Exit(1)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  scanner.Scan()
  numbers := strings.Split(scanner.Text(), ",")
  fmt.Println(numbers)

  var boards []Board
  hasMore := scanner.Scan()
  for hasMore {
    boards = append(boards, readBoard(scanner))
    hasMore = scanner.Scan()
  }

  fmt.Println(boards)
}
