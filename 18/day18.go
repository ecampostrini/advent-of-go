package main

import (
	"fmt"
	"strconv"
	//"unicode"
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

func parseNumber(in string) BinTree {
	if n, err := strconv.Atoi(in); err == nil {
		return BinTree{n, nil, nil, nil}
	}
	p1, p2 := partition(in)
	n1, n2 := parseNumber(p1), parseNumber(p2)
	ret := BinTree{0, &n1, &n2, nil}
	n1.Parent, n2.Parent = &ret, &ret
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
	if depth == 4 && node.Left != nil && node.Right != nil{
    fmt.Printf("Reached explosion depth: %v\n", node)
    fmt.Printf("Exploding: %v, %v\n", node.Left, node.Right)

		curr := node.Parent
    prev := node
		for curr != nil && (curr.Right == nil || curr.Right == prev){
      prev = curr
			curr = curr.Parent
		}
    fmt.Printf("current %v\n", curr)
		if curr != nil {
			curr = curr.Right
		}
		if firstRightLeaf := getFirstLeftLeaf(curr); firstRightLeaf != nil {
      fmt.Printf("firstRightLeaf: %v\n", firstRightLeaf)
      fmt.Printf("node.Right: %v\n", node.Right)
			firstRightLeaf.Val += node.Right.Val
			//fmt.Printf("firstRightLeafval: %v\n", firstRightLeaf.Val)
		}

		curr = node.Parent
    prev = node
		for curr != nil && (curr.Left == nil || curr.Left == prev){
      prev = curr
			curr = curr.Parent
		}
		if curr != nil {
			curr = curr.Left
		}
		if firstLeftLeaf := getFirstRightLeaf(curr); firstLeftLeaf != nil {
      fmt.Printf("firstLeftLeaf: %v\n", firstLeftLeaf)
      fmt.Printf("node.Left: %v\n", node.Left)
			firstLeftLeaf.Val += node.Left.Val
		}

    *node = BinTree{0, nil, nil, node.Parent}
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

func main() {
	//_, file := files.ReadFile("./input.txt")
	//defer file.Close()

	l0.Parent, l1.Parent = &n0, &n0
	n0.Parent, l2.Parent = &n1, &n1
	n1.Parent, l3.Parent = &n2, &n2
	n2.Parent, l2.Parent = &n3, &n3
	n3.Parent, l5.Parent = &n4, &n4
	//printBinTree(&n4)
	//explode(&n4, 0)
	//printBinTree(&n4)
	//p1, p2 := partition("[[[[4,3],[1,7]],[4,[9,2]]],[[6,[1,7]],[[8,0],3]]]")
	//fmt.Printf("%v - %v\n", p1, p2)
	n := parseNumber("[7,[6,[5,[4,[3,2]]]]]")
	printBinTree(&n)
  explode(&n, 0)
  printBinTree(&n)

}
