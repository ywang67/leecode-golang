package main

import (
	"encoding/json"
	"fmt"
)

type treeNode struct {
	Val       int
	LeftTree  *treeNode
	RightTree *treeNode
}

func main() {
	var root *treeNode
	root = insertNode(root, 7)
	root = insertNode(root, 11)
	root = insertNode(root, 3)
	root = insertNode(root, 8)
	root = insertNode(root, 12)
	root = insertNode(root, 9)
	root = insertNode(root, 6)

	test, _ := json.Marshal(root)
	fmt.Println("000tester: ", string(test))

	res := searchNode(root, 6)
	test2, _ := json.Marshal(res)
	fmt.Println("111tester: ", string(test2))
}

func insertNode(root *treeNode, value int) *treeNode {
	if root == nil {
		return &treeNode{
			Val: value,
		}
	}
	if root.Val == value {
		return root
	}

	if root.Val > value {
		root.LeftTree = insertNode(root.LeftTree, value)
	} else {
		root.RightTree = insertNode(root.RightTree, value)
	}

	return root
}

func searchNode(root *treeNode, value int) *treeNode {
	if root == nil {
		return nil
	}

	if root.Val == value {
		return root
	}

	if root.Val > value {
		return searchNode(root.LeftTree, value)
	} else {
		return searchNode(root.RightTree, value)
	}
}
