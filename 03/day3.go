package main

import (
  "fmt"
  "bufio"
  "os"
)

func readInput(filePath string)([]string, error){
  file, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var result []string
  for scanner.Scan() {
    result = append(result, scanner.Text())
  }

  return result, scanner.Err()
}

func boolFilter(input []string, pos int, filterSet bool)([]string) {
  var ret []string
  for _, num := range input {
    targetBit := num[pos]
    if (targetBit == '1' && filterSet) || (targetBit == '0' && !filterSet){
      ret = append(ret, num)
    }
  }
  return ret
}

func countBitsPerCol(input []string, row int) int {
  ret := 0
  for _, number := range input {
    ret += int(number[row]) - int('0')
  }
  return ret
}

func binary2Int(input string) int {
  ret := 0
  for i := 0; i < len(input); i++ {
    if input[i] == '1' {
      ret += 1 << (len(input) - 1 - i)
    }
  }
  return ret
}

func main() {
  input, err := readInput("./input.txt")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Oops")
    os.Exit(1)
  }

  numCount, numLenght := len(input), len(input[0])
  bitsSetPerRow := make([]int, numLenght)
  for _, number := range input {
    for idx, digit := range number { 
      bitsSetPerRow[idx] += int(digit) - int('0')
    }
  }

  gammaRate, epsilonRate, treshold := 0, 0, numCount / 2
  for idx, digitCount := range bitsSetPerRow {
    pos := (numLenght - 1) - idx
    if digitCount >= treshold { 
      gammaRate = gammaRate | 1 << pos
    } else {
      epsilonRate = epsilonRate | 1 << pos
    }
  }
  fmt.Println(gammaRate, epsilonRate, gammaRate * epsilonRate)

  oxigenGeneratorRating, co2ScrubberRating := input, input
  for i := 0; i < numLenght; i++ {
    if len(oxigenGeneratorRating) > 1 {
      oxigenGeneratorFlag := 2 * countBitsPerCol(oxigenGeneratorRating, i) >= len(oxigenGeneratorRating)
      oxigenGeneratorRating = boolFilter(oxigenGeneratorRating, i, oxigenGeneratorFlag)
    }

    if len(co2ScrubberRating) > 1 {
      co2ScrubberRatingFlag := 2 * countBitsPerCol(co2ScrubberRating, i) < len(co2ScrubberRating)
      co2ScrubberRating = boolFilter(co2ScrubberRating, i, co2ScrubberRatingFlag)
    }
  }

  oxigenGeneratorRatingDec := binary2Int(oxigenGeneratorRating[0])
  co2ScrubberRatingDec :=  binary2Int(co2ScrubberRating[0])
  fmt.Println(oxigenGeneratorRatingDec, 
    co2ScrubberRatingDec, 
    oxigenGeneratorRatingDec * co2ScrubberRatingDec)

}
