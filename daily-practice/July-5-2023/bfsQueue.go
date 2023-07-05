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

	fmt.Println("000tester: ", bfsQueque(nodeA))
}

func bfsQueque(root *Node) []string {
	if root == nil {
		return nil
	}

	dp := []*Node{root}
	visited := make(map[*Node]bool)

	res := []string{}

	for len(dp) > 0 {
		tempNode := dp[0]
		dp = dp[1:]
		dp = append(dp, tempNode.Children...)
		if !visited[tempNode] {
			res = append(res, tempNode.Val)
		}
	}

	return res
}
