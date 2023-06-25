package main

import (
	"encoding/json"
	"fmt"
)

// please create a binary tree with insertNode and searchNode

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		root.LeftNode = insertNode(root.LeftNode, val)
	} else {
		root.RightNode = insertNode(root.RightNode, val)
	}

	return root
}

func searchNode(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if val < root.Val {
		return searchNode(root.LeftNode, val)
	}

	return searchNode(root.RightNode, val)
}

func main() {
	root := &TreeNode{
		Val: 4,
	}

	insertNode(root, 2)
	insertNode(root, 1)

	test, _ := json.Marshal(root)
	fmt.Println("inserted root: ", string(test))

	res1 := searchNode(root, 2)
	test1, _ := json.Marshal(res1)
	fmt.Println("search root: ", string(test1))

	res2 := searchNode(root, 5)
	test2, _ := json.Marshal(res2)
	fmt.Println("search root: ", string(test2))
}
