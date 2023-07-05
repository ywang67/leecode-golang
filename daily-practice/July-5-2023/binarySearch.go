package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func insertNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value}
	}

	if root.Value == value {
		return root
	}

	if root.Value > value {
		root.LeftNode = insertNode(root.LeftNode, value)
	} else {
		root.RightNode = insertNode(root.RightNode, value)
	}

	return root
}

func searchNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Value > value {
		return searchNode(root.LeftNode, value)
	}
	if root.Value < value {
		return searchNode(root.RightNode, value)
	}

	return root
}

func main() {
	root := insertNode(nil, 9)
	root = insertNode(root, 7)
	root = insertNode(root, 3)
	root = insertNode(root, 1)
	root = insertNode(root, 8)
	root = insertNode(root, 5)
	root = insertNode(root, 2)
	root = insertNode(root, 7)
	root = insertNode(root, 11)

	// test, _ := json.Marshal(root)
	// fmt.Println("000tester: ", string(test))

	test2, _ := json.Marshal(searchNode(root, 1))
	fmt.Println("111tester: ", string(test2))
}
