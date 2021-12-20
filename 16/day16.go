package main

import (
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"math"
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
	for readLast := false; !readLast; {
		if input[*pos] == '0' {
			readLast = true
		}
		*pos += 1
		acc += input[*pos : (*pos)+4]
		*pos += 4
	}
	return binString2Int(acc)
}

func parsePacket(input string, pos *int) Packet {
	version, typeId := parseHeader(input, pos)
	if typeId == 4 {
		value := parseValue(input, pos)
		return Packet{version, typeId, []*Packet{}, value}
	}

	lengthTypeId := string(input[*pos])
	*pos++
	if lengthTypeId == "0" {
		subpacketsLength := binString2Int(input[*pos : *pos+15])
		*pos += 15
		var subpackets []*Packet
		var totalRead int
		for totalRead != subpacketsLength {
			currentRead := 0
			newPacket := parsePacket(input[*pos+totalRead:], &currentRead)
			subpackets = append(subpackets, &newPacket)
			totalRead += currentRead
		}
		*pos += subpacketsLength
		return Packet{version, typeId, subpackets, -1}
	} else if lengthTypeId == "1" {
		subpacketsCount := binString2Int(input[*pos : *pos+11])
		*pos += 11
		var subpackets []*Packet
		var totalRead int
		for subpacketsCount > 0 {
			currentRead := 0
			newPacket := parsePacket(input[*pos+totalRead:], &currentRead)
			subpackets = append(subpackets, &newPacket)
			totalRead += currentRead
			subpacketsCount--
		}
		*pos += totalRead
		return Packet{version, typeId, subpackets, -1}
	} else {
		panic(fmt.Sprintf("lengthTypeId not supported: %d", lengthTypeId))
	}
}

func printPacket(packet *Packet, indent string) {
	if packet.TypeId == 4 {
		fmt.Printf("%s%d, %d, %d\n", indent, packet.Version, packet.TypeId, packet.Value)
	} else {
		fmt.Printf("%s%d, %d:\n", indent, packet.Version, packet.TypeId)
		for _, p := range packet.Children {
			printPacket(p, indent+"|-")
		}
	}
}

func countVersions(packet *Packet) int {
	ret := packet.Version
	for _, p := range packet.Children {
		ret += countVersions(p)
	}
	return ret
}

func eval(packet *Packet) int {
	typeId := packet.TypeId
	if typeId == 0 {
		var sum int
		for _, p := range packet.Children {
			sum += eval(p)
		}
		return sum
	} else if typeId == 1 {
		prod := 1
		for _, p := range packet.Children {
			prod *= eval(p)
		}
		return prod
	} else if typeId == 2 {
		min := math.MaxInt
		for _, p := range packet.Children {
			val := eval(p)
			if val < min {
				min = val
			}
		}
		return min
	} else if typeId == 3 {
		max := math.MinInt
		for _, p := range packet.Children {
			val := eval(p)
			if val > max {
				max = val
			}
		}
		return max
	} else if typeId == 4 {
		return packet.Value
	} else if typeId == 5 {
		v1, v2 := eval(packet.Children[0]), eval(packet.Children[1])
		if v1 > v2 {
			return 1
		} else {
			return 0
		}
	} else if typeId == 6 {
		v1, v2 := eval(packet.Children[0]), eval(packet.Children[1])
		if v1 < v2 {
			return 1
		} else {
			return 0
		}
	} else if typeId == 7 {
		v1, v2 := eval(packet.Children[0]), eval(packet.Children[1])
		if v1 == v2 {
			return 1
		} else {
			return 0
		}
	} else {
		panic(fmt.Sprintf("TypeId not supported: %d", typeId))
	}
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()
	pos := 0
	scanner.Scan()
	mainPacket := parsePacket(hexString2BinString(scanner.Text()), &pos)

	// for debugging purposes
	//printPacket(&mainPacket, "")
	// part 1
	fmt.Println(countVersions(&mainPacket))
	// part 2
	fmt.Println(eval(&mainPacket))
}
