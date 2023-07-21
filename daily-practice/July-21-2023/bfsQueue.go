package main

import (
	"fmt"
)

type Node struct {
	Value string
	Children []*Node
}

func main() {
	nodeA := &Node{Value: "a"}
	nodeB := &Node{Value: "b"}
	nodeC := &Node{Value: "c"}
	nodeD := &Node{Value: "d"}
	nodeE := &Node{Value: "e"}
	nodeF := &Node{Value: "f"}
	nodeG := &Node{Value: "g"}

	nodeA.Children = append(nodeA.Children, nodeB, nodeC)
	nodeB.Children = append(nodeB.Children, nodeD, nodeE)
	nodeC.Children = append(nodeC.Children, nodeF, nodeG)

	res := []string{}
	fmt.Println("000tester: ", bfsQueue(nodeA, res))
}

func bfsQueue(root *Node, res []string) []string {
	if root == nil {
		return res
	}

	queue := []*Node{root}
	isVisited := make(map[*Node]bool)

	for len(queue) > 0 {
		tempNode := queue[0]
		if !isVisited[tempNode] {
			queue = queue[1:]
			res = append(res, tempNode.Value)
			for _, child := range tempNode.Children {
				queue = append(queue, child)
			}
		}
	}

	return res
}
