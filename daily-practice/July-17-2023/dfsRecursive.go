package main

import "fmt"

type Node struct {
	Val string
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
	nodeH := &Node{Val: "h"}
	nodeI := &Node{Val: "i"}
	
	nodeA.Children = append(nodeA.Children, nodeB, nodeC)
	nodeB.Children = append(nodeC.Children, nodeD, nodeE)
	nodeC.Children = append(nodeE.Children, nodeF, nodeG)
	nodeD.Children = append(nodeD.Children, nodeH, nodeI)

	fmt.Println("000tester: ", dfsRecusive(nodeA))
}

func dfsRecusive(node *Node) []string {

	res := []string{}

	var dfsHelper func(root *Node)

	dfsHelper = func(root *Node) {
		if node == nil {
			return 
		}
		res = append(res, root.Val)
	
		for _, child := range root.Children {
			dfsHelper(child)
		}
	}

	dfsHelper(node)

	return res
}
