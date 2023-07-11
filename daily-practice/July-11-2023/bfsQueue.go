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

	fmt.Println("000tester: ", bfsQueue(rootA))
}

func bfsQueue(root *Node) []string {
	if root == nil {
		return nil
	}

	if len(root.Children) == 0 {
		return []string{root.Val}
	}

	queue := []*Node{root}
	visited := make(map[*Node]bool)

	res := []string{}
	for len(queue) > 0 {
		tempNode := queue[0]
		queue = queue[1:]
		if !visited[tempNode] {
			res = append(res, tempNode.Val)
			for _, child := range tempNode.Children {
				queue = append(queue, child)
			}
		}
	}

	return res
}
