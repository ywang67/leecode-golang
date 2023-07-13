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

func main() {
	var root *TreeNode
	root = insertNode(root, 9)
	root = insertNode(root, 2)
	root = insertNode(root, 3)
	root = insertNode(root, 7)
	root = insertNode(root, 11)
	root = insertNode(root, 12)
	root = insertNode(root, 17)

	test, _ := json.Marshal(root)
	fmt.Println("000tester: ", string(test))

	res := searchNode(root, 17)
	test2, _ := json.Marshal(res)
	fmt.Println("111tester: ", string(test2))
}

func insertNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Value: value,
		}
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
	if root == nil || root.Value == value {
		return root
	}

	if root.Value < value {
		return searchNode(root.RightNode, value)
	} else {
		return searchNode(root.LeftNode, value)
	}
}
