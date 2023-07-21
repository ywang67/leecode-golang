package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val int
	LeftNode *TreeNode
	RightNode *TreeNode
}

func main() {
	var root TreeNode
	insertNode(&root, 9)
	insertNode(&root, 11)
	insertNode(&root, 7)
	insertNode(&root, 8)
	insertNode(&root, 3)
	insertNode(&root, 5)
	insertNode(&root, 6)
	insertNode(&root, 17)

	test, _ := json.Marshal(root)
	fmt.Println("000tester: ", string(test))

	res := searchNode(&root, 17)
	test2, _ := json.Marshal(res)
	fmt.Println("111tester: ", string(test2))
}

func insertNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		node = &TreeNode{Val: val}
		return node
	}
	
	if node.Val > val {
		node.LeftNode = insertNode(node.LeftNode, val)
	} else {
		node.RightNode =  insertNode(node.RightNode, val)
	}

	return node
}

func searchNode(node *TreeNode, val int) *TreeNode {
	if node == nil || node.Val == val {
		return node
	}

	if node.Val > val {
		return searchNode(node.LeftNode, val)
	} else {
		return searchNode(node.RightNode, val)
	}
}
