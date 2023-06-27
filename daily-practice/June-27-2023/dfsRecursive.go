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

	res := []string{}
	fmt.Println("000tester: ", dfsRecursive(rootA, res))
}

func dfsRecursive(root *Node, res []string) []string {
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	for _, child := range root.Children {
		res = dfsRecursive(child, res)
	}
	return res
}
