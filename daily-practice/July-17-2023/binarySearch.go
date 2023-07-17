package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Value int
	LeftTree *TreeNode
	RightTree *TreeNode
}

func main() {
	arr := []int{7, 2, 9, 3, 5, 1, 11, 15}
	var root *TreeNode
	for _, arrEle := range arr {
		root = InsertNode(root, arrEle)
	}

	test, _ := json.Marshal(root)
	fmt.Println("000tester: ", string(test))

	test2, _ := json.Marshal(searchNode(root, 15))
	fmt.Println("111tester: ", string(test2))
}

func InsertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Value: val,
		}
	}
	if root.Value == val {
		return root
	}

	if root.Value > val {
		root.LeftTree =InsertNode(root.LeftTree, val)
	} else {
		root.RightTree = InsertNode(root.RightTree, val)
	}

	return root
}

func searchNode(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Value == val {
		return root
	}

	if root.Value > val {
		return searchNode(root.LeftTree, val)
	} else {
		return searchNode(root.RightTree, val)
	}
}
