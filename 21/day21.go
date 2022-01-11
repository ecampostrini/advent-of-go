package main

import (
	"fmt"
	//"github.com/ecampostrini/advent-of-go/utils/files"
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

func main() {
	//scanner, file := files.ReadFile("./input.test.txt")
	//defer file.Close()

	dice := getDeterministicDice()
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d: %d\n", i, dice())
	}
}
