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
	for readLast := false; !readLast; {
		if input[*pos] == '0' {
			readLast = true
		}
		*pos += 1
		acc += input[*pos : (*pos)+4]
		*pos += 4
	}
	//if (*pos % 4) != 0 {
	//upperBound := 4 * (*pos/4 + 1)
	//*pos += upperBound - *pos
	//}
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
	} else {
		panic("Length type id not supported!")
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

func main() {
	scanner, file := files.ReadFile("./lengthTypeId0.test.txt")
	defer file.Close()
	pos := 0
	scanner.Scan()
	mainPacket := parsePacket(hexString2BinString(scanner.Text()), &pos)
  printPacket(&mainPacket, "")
	//fmt.Printf("%v\n", mainPackage)
}
