package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func readNumbers(filePath string)([]int, error){
  file, _ := os.Open(filePath)
  scanner := bufio.NewScanner(file)
  var result []int
  for scanner.Scan(){
    x, err := strconv.Atoi(scanner.Text())
    if err != nil {
      return result, err
    }
    result = append(result, x)
  }
  return result, scanner.Err()
}

func main() {
  nums, _ := readNumbers("./input.txt")

  //part 1
  count := 0
  for i := 1; i < len(nums); i++ {
    if nums[i-1] < nums[i] {
      count++
    }
  }
  fmt.Println("Part1:", count)

  //part 2
  var currWindow int;
  count = 0
  for i := 0; i < len(nums); i++ {
    if i < 3 {
      currWindow += nums[i]
    } else {
      nextWindow := currWindow - nums[i-3] + nums[i]
      if nextWindow > currWindow {
        count++
      }
    }
  }
  fmt.Println("Part2:", count)
}
