package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func initLanternFish(s *bufio.Scanner) map[int]int {
  ret := make(map[int]int)

  s.Scan()
  for _, c := range strings.Split(s.Text(), ",") {
    fmt.Println(c) 
    n, err := strconv.Atoi(c)
    if err != nil {
      fmt.Println("Failed to parse number from input", err)
      os.Exit(1)
    }
    ret[n]++
  }
  return ret
}

func main() {
  file, err := os.Open("./input.test.txt")
  if err != nil {
    fmt.Println("Error opening input file", err)
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  lanternfish := initLanternFish(scanner)

  fmt.Println(lanternfish)
}
