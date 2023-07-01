package main

import "fmt"

type Node struct {
	Val      string
	Children []*Node
}

func main() {
	nodeA := Node{Val: "a"}
	nodeB := Node{Val: "b"}
	nodeC := Node{Val: "c"}
	nodeD := Node{Val: "d"}
	nodeE := Node{Val: "e"}
	nodeF := Node{Val: "f"}
	nodeG := Node{Val: "g"}

	nodeA.Children = append(nodeA.Children, &nodeB, &nodeC)
	nodeB.Children = append(nodeB.Children, &nodeD, &nodeE)
	nodeC.Children = append(nodeC.Children, &nodeF, &nodeG)

	bfsQueue(&nodeA)
}

func bfsQueue(root *Node) {
	queue := []*Node{root}
	visited := make(map[*Node]bool)

	for len(queue) > 0 {
		tempNode := queue[0]
		queue = queue[1:]

		if !visited[tempNode] {
			visited[tempNode] = true

			fmt.Println("000tester: ", tempNode.Val)
			for _, child := range tempNode.Children {
				queue = append(queue, child)
			}
		}
	}
}
