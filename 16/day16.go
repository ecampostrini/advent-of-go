package main

import (
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
)

type Packet struct {
	Version  int
	TypeId   int
	Children []*Packet
	Value    int
}

var hexConv = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func hexString2BinString(input string) string {
	var ret string
	for _, c := range input {
		ret += hexConv[string(c)]
	}
	return ret
}

func binString2Int(input string) int {
	var ret int
	for _, c := range input {
		ret = ret << 1
		if string(c) == "1" {
			ret = ret | 1
		}
	}
	return ret
}

func parseHeader(input string, pos *int) (int, int) {
	version := input[*pos : (*pos)+3]
	(*pos) += 3
	typeId := input[*pos : (*pos)+3]
	(*pos) += 3
	return binString2Int(version), binString2Int(typeId)
}

func parseValue(input string, pos *int) int {
	var acc string
	fmt.Println(input)
	for input[*pos] != 0 {
		fmt.Println("Pos", *pos)
		fmt.Println("ip", string(input[*pos]))
		fmt.Println("Acc", acc)
		*pos += 1
		acc += input[*pos : (*pos)+4]
		*pos += 4
	}
	fmt.Println("Acc", acc)
	//*pos += 1
	//acc += input[*pos : (*pos)+4]
	//*pos += 4
	return binString2Int(acc)
}

func parsePacket(input string, pos *int) Packet {
	fmt.Println("Pos", *pos)
	version, typeId := parseHeader(input, pos)
	fmt.Println("Pos", *pos)
	if typeId == 4 {
		fmt.Println("here")
		value := parseValue(input, pos)
		fmt.Println("there")
		return Packet{version, typeId, []*Packet{}, value}
	} else {
		panic("Version not supported yet!")
	}
}

func main() {
	scanner, file := files.ReadFile("./input.test.txt")
	defer file.Close()
	pos := 0
	scanner.Scan()
	mainPackage := parsePacket(hexString2BinString(scanner.Text()), &pos)
	fmt.Printf("%v\n", mainPackage)
}
