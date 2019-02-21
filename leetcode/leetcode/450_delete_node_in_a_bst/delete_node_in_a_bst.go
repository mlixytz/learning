package leetcode450

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	} else {

	}
	m := min(root.Right)
	m.Right = deleteMin(root.Right)
	m.Left = root.Left
	return m
}

func min(node *TreeNode) *TreeNode {
	for node.Left == nil {
		return node
	}
	return min(node.Left)
}

func deleteMin(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteMin(node.Left)
	return node
}
