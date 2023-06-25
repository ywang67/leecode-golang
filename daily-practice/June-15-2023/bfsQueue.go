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

	bfsQueue(nodeA)
}

func bfsQueue(node *Node) {
	if node == nil {
		return
	}

	visited := make(map[*Node]bool)
	queue := []*Node{node}

	for len(queue) > 0 {
		tempNode := queue[0]
		queue = queue[1:]
		// fmt.Println("111tester: ", tempNode.Val)
		if !visited[tempNode] {
			fmt.Println("000tester: ", tempNode.Val)
			visited[tempNode] = true
			for _, child := range tempNode.Children {
				queue = append(queue, child)
			}
		}
	}
}
