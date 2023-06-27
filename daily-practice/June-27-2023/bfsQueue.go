package main

import "fmt"

type Node struct {
	Val      string
	Children []*Node
}

func main() {
	rootA := &Node{
		Val: "a",
	}
	rootB := &Node{
		Val: "b",
	}
	rootC := &Node{
		Val: "c",
	}
	rootD := &Node{
		Val: "d",
	}
	rootE := &Node{
		Val: "e",
	}
	rootF := &Node{
		Val: "f",
	}
	rootG := &Node{
		Val: "g",
	}

	rootA.Children = append(rootA.Children, rootB, rootC)
	rootB.Children = append(rootB.Children, rootD, rootE)
	rootC.Children = append(rootC.Children, rootF, rootG)

	fmt.Println("000tester: ", bfsQueue(rootA))
}

func bfsQueue(root *Node) []string {
	if root == nil {
		return nil
	}

	res := []string{}

	visited := make(map[*Node]bool)
	queue := []*Node{root}

	for len(queue) > 0 {
		tempNode := queue[0]
		res = append(res, tempNode.Val)
		queue = queue[1:]
		visited[tempNode] = true

		for _, node := range tempNode.Children {
			if !visited[node] {
				visited[node] = true
				queue = append(queue, node)
			}
		}
	}

	return res
}
