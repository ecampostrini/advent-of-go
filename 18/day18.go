package main

import (
	"fmt"
  "unicode"
	//"github.com/ecampostrini/advent-of-go/utils/files"
)

type BinTree struct {
	Val                 int
	Left, Right, Parent *BinTree
}

var l0 = BinTree{9, nil, nil, nil}
var l1 = BinTree{8, nil, nil, nil}
var n0 = BinTree{0, &l0, &l1, nil}

//l0.Parent, l1.Parent = &n0, &n0

var l2 = BinTree{1, nil, nil, nil}
var n1 = BinTree{0, &n0, &l2, nil}

//n0.Parent, l2.Parent = &n1, &n1

var l3 = BinTree{2, nil, nil, nil}
var n2 = BinTree{0, &n1, &l3, nil}

//n1.Parent, l3.Parent = &n2, &n2

var l4 = BinTree{3, nil, nil, nil}
var n3 = BinTree{0, &n2, &l4, nil}

//n2.Parent, l2.Parent = &n3, &n3

var l5 = BinTree{4, nil, nil, nil}
var n4 = BinTree{0, &n3, &l5, nil}

//n3.Parent, l5.Parent = &n4, &n4

func ParseTree(input string) BinTree {
}

func PrintBinTree(r *BinTree) {
	if r.Left == nil && r.Right == nil {
		fmt.Printf("%d", r.Val)
		return
	}
	fmt.Printf("[")
	if r.Left != nil {
		PrintBinTree(r.Left)
	}
	fmt.Printf(",")
	if r.Right != nil {
		PrintBinTree(r.Right)
	}
	fmt.Printf("]")
}

func getFirstLeftLeaf(node *BinTree) *BinTree {
	if node.Left == nil && node.Right == nil {
		return node
	}
	return getFirstLeftLeaf(node.Left)
}

func getFirstRightLeaf(node *BinTree) *BinTree {
	if node.Left == nil && node.Right == nil {
		return node
	}
	return getFirstRightLeaf(node.Right)
}

func Explode(node *BinTree, depth int) bool {
	if depth == 4 {
		fmt.Printf("Exploding: %v, %v\n", node.Right, node.Left)

		curr := node.Parent
		for curr != nil && curr.Right == nil {
			curr = curr.Parent
		}
		if curr != nil {
			curr = curr.Right
		}
		if firstRightLeaf := getFirstLeftLeaf(curr); firstRightLeaf != nil {
			fmt.Printf("firstRightLeaf: %v\n", firstRightLeaf)
			firstRightLeaf.Val += node.Right.Val
		}

		curr = node.Parent
		for curr != nil && curr.Left == nil {
			curr = curr.Parent
		}
		if curr != nil {
			curr = curr.Left
		}
		if firstLeftLeaf := getFirstRightLeaf(curr); firstLeftLeaf != nil {
			firstLeftLeaf.Val += node.Left.Val
		}

		if node.Parent.Left == node {
			node.Parent.Left = &BinTree{0, nil, nil, nil}
		} else {
			node.Parent.Right = &BinTree{0, nil, nil, nil}
		}
		return true
	}

	var hadExplosion bool
	if node.Left != nil {
		hadExplosion = Explode(node.Left, depth+1)
	}
	if !hadExplosion && node.Right != nil {
		hadExplosion = Explode(node.Right, depth+1)
	}
	return hadExplosion
}

func main() {
	//_, file := files.ReadFile("./input.txt")
	//defer file.Close()

	l0.Parent, l1.Parent = &n0, &n0
	n0.Parent, l2.Parent = &n1, &n1
	n1.Parent, l3.Parent = &n2, &n2
	n2.Parent, l2.Parent = &n3, &n3
	n3.Parent, l5.Parent = &n4, &n4
	PrintBinTree(&n4)
	Explode(&n4, 0)
	PrintBinTree(&n4)

}
