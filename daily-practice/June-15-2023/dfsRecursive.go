package main

import "fmt"

type Node struct {
	Val      string
	Children []*Node
}

func main() {
	nodeA := &Node{
		Val: "A",
	}
	nodeB := &Node{
		Val: "B",
	}
	nodeC := &Node{
		Val: "C",
	}
	nodeD := &Node{
		Val: "D",
	}
	nodeE := &Node{
		Val: "E",
	}
	nodeF := &Node{
		Val: "F",
	}
	nodeG := &Node{
		Val: "G",
	}
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

	for _, child := range root.Children {
		dfsRecursive(child)
	}
}
