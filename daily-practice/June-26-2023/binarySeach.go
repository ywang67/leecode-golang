package main

import (
	"fmt"
)

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		root.LeftNode = &TreeNode{Val: val}
	} else {
		root.RightNode = &TreeNode{Val: val}
	}

	return root
}

func searchNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val < val {
		return searchNode(root.LeftNode, val)
	}

	return searchNode(root.RightNode, val)
}

func main() {
	var root *TreeNode
	root = insertNode(root, 9)
	insertNode(root, 3)
	insertNode(root, 15)
	insertNode(root, 7)
	insertNode(root, 11)
	insertNode(root, 6)

	fmt.Println("000tester: ", searchNode(root, 9))
}
