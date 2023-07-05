package main

import "fmt"

type Node struct {
	Val      string
	Children []*Node
}

func main() {
	nodeA := &Node{Val: "a"}
	nodeB := &Node{Val: "b"}
	nodeC := &Node{Val: "c"}
	nodeD := &Node{Val: "d"}
	nodeE := &Node{Val: "e"}
	nodeF := &Node{Val: "f"}
	nodeG := &Node{Val: "g"}

	nodeA.Children = append(nodeA.Children, nodeB, nodeC)
	nodeB.Children = append(nodeB.Children, nodeD, nodeE)
	nodeC.Children = append(nodeC.Children, nodeF, nodeG)

	dfsRecursive(nodeA)
}

func dfsRecursive(root *Node) {
	if root == nil {
		return
	}

	fmt.Println("000tester: ", root.Val)

	for _, node := range root.Children {
		dfsRecursive(node)
	}
}
