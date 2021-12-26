package main

import (
	"fmt"
	"github.com/ecampostrini/advent-of-go/utils/files"
	"math"
	"strconv"
)

type BinTree struct {
	Val                 int
	Left, Right, Parent *BinTree
}

func partition(in string) (string, string) {
	var intStack, partitionPos int
	for pos, c := range in {
		if c == '[' {
			intStack++
		} else if c == ']' {
			intStack--
		} else if c == ',' && intStack == 1 {
			partitionPos = pos
			break
		}
	}
	return in[1:partitionPos], in[partitionPos+1 : len(in)-1]
}

func parseNumber(in string) *BinTree {
	if n, err := strconv.Atoi(in); err == nil {
		return &BinTree{n, nil, nil, nil}
	}
	p1, p2 := partition(in)
	n1, n2 := parseNumber(p1), parseNumber(p2)
	ret := &BinTree{0, n1, n2, nil}
	n1.Parent, n2.Parent = ret, ret
	return ret
}

func printBinTree(r *BinTree) {
	if r.Left == nil && r.Right == nil {
		fmt.Printf("%d", r.Val)
		return
	}
	fmt.Printf("[")
	if r.Left != nil {
		printBinTree(r.Left)
	}
	fmt.Printf(",")
	if r.Right != nil {
		printBinTree(r.Right)
	}
	fmt.Printf("]")
}

func getFirstLeftLeaf(node *BinTree) *BinTree {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return node
	}
	return getFirstLeftLeaf(node.Left)
}

func getFirstRightLeaf(node *BinTree) *BinTree {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return node
	}
	return getFirstRightLeaf(node.Right)
}

func explode(node *BinTree, depth int) bool {
	if depth == 4 && (node.Left != nil || node.Right != nil) {
		curr := node.Parent
		prev := node
		for curr != nil && (curr.Right == nil || curr.Right == prev) {
			prev = curr
			curr = curr.Parent
		}
		if curr != nil {
			curr = curr.Right
		}
		if firstRightLeaf := getFirstLeftLeaf(curr); firstRightLeaf != nil {
			firstRightLeaf.Val += node.Right.Val
		}

		curr = node.Parent
		prev = node
		for curr != nil && (curr.Left == nil || curr.Left == prev) {
			prev = curr
			curr = curr.Parent
		}
		if curr != nil {
			curr = curr.Left
		}
		if firstLeftLeaf := getFirstRightLeaf(curr); firstLeftLeaf != nil {
			firstLeftLeaf.Val += node.Left.Val
		}
		node.Left = nil
		node.Right = nil
		node.Val = 0
		return true
	}

	var hadExplosion bool
	if node.Left != nil {
		hadExplosion = explode(node.Left, depth+1)
	}
	if !hadExplosion && node.Right != nil {
		hadExplosion = explode(node.Right, depth+1)
	}
	return hadExplosion
}

func split(node *BinTree) bool {
	if node == nil {
		return false
	}
	if node.Val >= 10 {
		half := node.Val / 2
		a := &BinTree{half, nil, nil, nil}
		b := &BinTree{half + node.Val%2, nil, nil, nil}
		c := &BinTree{0, a, b, node.Parent}
		a.Parent, b.Parent = c, c
		if node.Parent.Left == node {
			node.Parent.Left = c
		} else {
			node.Parent.Right = c
		}
		return true
	}
	var hadSplit bool
	hadSplit = split(node.Left)
	if !hadSplit {
		hadSplit = split(node.Right)
	}
	return hadSplit
}

func reduce(node *BinTree) {
	hadReduction := true
	for hadReduction {
		hadExplosion := false
		for explode(node, 0) {
			hadExplosion = true
		}
		hadSplit := split(node)
		hadReduction = hadExplosion || hadSplit
	}
}

func getMagnitude(node *BinTree) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return node.Val
	}
	return 3*getMagnitude(node.Left) + 2*getMagnitude(node.Right)
}

func main() {
	scanner, file := files.ReadFile("./input.txt")
	defer file.Close()

	var snailfishNumbers []string
	for scanner.Scan() {
		snailfishNumbers = append(snailfishNumbers, scanner.Text())
	}

	result := parseNumber(snailfishNumbers[0])
	for i := 1; i < len(snailfishNumbers); i++ {
		newNumber := parseNumber(snailfishNumbers[i])
		newPair := &BinTree{0, result, newNumber, nil}
		result.Parent, newNumber.Parent = newPair, newPair
		result = newPair
		reduce(result)
	}
	// for debugging
	//printBinTree(result)
	//fmt.Printf("\n")

	// part 1
	fmt.Println(getMagnitude(result))

	// part 2
	var largestMagnitude = math.MinInt
	for i := 0; i < len(snailfishNumbers)-1; i++ {
		var root *BinTree
		for j := i; j < len(snailfishNumbers); j++ {
			n1, n2 := parseNumber(snailfishNumbers[i]), parseNumber(snailfishNumbers[j])
			root = &BinTree{0, n1, n2, nil}
			n1.Parent, n2.Parent = root, root
			reduce(root)
			if magnitude := getMagnitude(root); magnitude > largestMagnitude {
				largestMagnitude = magnitude
			}

			n1, n2 = parseNumber(snailfishNumbers[i]), parseNumber(snailfishNumbers[j])
			root = &BinTree{0, n2, n1, nil}
			n1.Parent, n2.Parent = root, root
			reduce(root)
			if magnitude := getMagnitude(root); magnitude > largestMagnitude {
				largestMagnitude = magnitude
			}
		}
	}
	fmt.Println(largestMagnitude)
}
