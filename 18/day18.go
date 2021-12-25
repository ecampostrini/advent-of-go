package main

import (
	"fmt"
	"strconv"
	//"github.com/ecampostrini/advent-of-go/utils/files"
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
	if depth == 4 && node.Left != nil && node.Right != nil {
		//fmt.Printf("Reached explosion depth: %v\n", node)
		//fmt.Printf("Exploding: %v, %v\n", node.Left, node.Right)

		curr := node.Parent
		prev := node
		for curr != nil && (curr.Right == nil || curr.Right == prev) {
			//fmt.Printf("current %v\n", curr)
			prev = curr
			curr = curr.Parent
		}
		//fmt.Printf("current %v\n", curr)
		if curr != nil {
			curr = curr.Right
		}
		if firstRightLeaf := getFirstLeftLeaf(curr); firstRightLeaf != nil {
			//fmt.Printf("firstRightLeaf: %v\n", firstRightLeaf)
			//fmt.Printf("node.Right: %v\n", node.Right)
			firstRightLeaf.Val += node.Right.Val
			//fmt.Printf("firstRightLeafval: %v\n", firstRightLeaf.Val)
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
			//fmt.Printf("firstLeftLeaf: %v\n", firstLeftLeaf)
			//fmt.Printf("node.Left: %v\n", node.Left)
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

func split(node *BinTree) bool {
	if node == nil {
		return false
	}

	if node.Val >= 10 {
		half := node.Val / 2
		a := &BinTree{half, nil, nil, node}
		b := &BinTree{half + node.Val%2, nil, nil, node}
		*node = BinTree{0, a, b, node.Parent}
		return true
	}

	var hadSplit bool
	hadSplit = split(node.Left)
	if !hadSplit {
		hadSplit = split(node.Right)
	}
	return hadSplit
}

func inOrder(n *BinTree) {
	if n == nil {
		return
	}

	inOrder(n.Left)
	fmt.Printf("(%p) %v\n", n, *n)
	inOrder(n.Right)
}

func printAndExplode(in string) {
	n := parseNumber(in)
	//fmt.Println(in)
	//inOrder(n)
	fmt.Printf("\n---\n")
	printBinTree(n)
	explode(n, 0)
	fmt.Printf("\n")
	printBinTree(n)
	explode(n, 0)
	fmt.Printf("\n")
	printBinTree(n)
	explode(n, 0)
	fmt.Printf("\n")
	printBinTree(n)
}

func main() {
	//_, file := files.ReadFile("./input.txt")
	//defer file.Close()

	var testInput = []string{
		"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
		"[[[[[9,8],1],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[[[4,3],[1,7]],[4,[9,2]]],[[6,[1,7]],[[8,0],3]]]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
	}

	for _, in := range testInput {
		n := parseNumber(in)
		hadReduction := true
		fmt.Printf("\n---\n")

		for hadReduction {
			fmt.Printf("\n")
			printBinTree(n)

			hadExplosion := false
			for explode(n, 0) {
				hadExplosion = true
				fmt.Printf("\n")
				//printBinTree(n)
			}

			hadSplit := false
			for split(n) {
				hadSplit = true
				fmt.Printf("\n")
				//printBinTree(n)
			}

			hadReduction = hadExplosion || hadSplit
		}
	}
}
