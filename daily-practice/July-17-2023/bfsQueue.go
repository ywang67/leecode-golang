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

	fmt.Println("000tester: ", bfsQueue(nodeA))
}

func bfsQueue(node *Node) []string {
	queue := []*Node{node}
	isVisited := make(map[*Node]bool)

	res  := []string{}
	for len(queue) > 0 {
		temp := queue[0]
		if !isVisited[temp] {
			queue = queue[1:]
			res = append(res, temp.Val)
			isVisited[temp] = true
			queue = append(queue, temp.Children...)
		}
	}

	return res
}
