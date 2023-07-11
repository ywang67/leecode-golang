package main

import "fmt"

type Node struct {
	Val      string
	Children []*Node
}

func main() {
	rootA := &Node{Val: "a"}
	rootB := &Node{Val: "b"}
	rootC := &Node{Val: "c"}
	rootD := &Node{Val: "d"}
	rootE := &Node{Val: "e"}
	rootF := &Node{Val: "f"}
	rootG := &Node{Val: "g"}

	rootA.Children = append(rootA.Children, rootB, rootC)
	rootB.Children = append(rootB.Children, rootD, rootE)
	rootC.Children = append(rootC.Children, rootF, rootG)

	fmt.Println("000tester: ", dfsRecursive(rootA))
}

func dfsRecursive(root *Node) []string {
	res := []string{}

	var dfsHelper func(node *Node)
	dfsHelper = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			dfsHelper(child)
		}
	}

	dfsHelper(root)

	return res
}
