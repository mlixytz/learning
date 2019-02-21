package leetcode623

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		newRoot := &TreeNode{Val: v, Left: root}
		return newRoot
	} else if d == 2 {
		root.Left = &TreeNode{Val: v, Left: root.Left}
		root.Right = &TreeNode{Val: v, Right: root.Right}
		return root
	}
	if root.Left != nil {
		root.Left = addOneRow(root.Left, v, d-1)
	}
	if root.Right != nil {
		root.Right = addOneRow(root.Right, v, d-1)
	}
	return root
}
