package leetcode701

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var prev *TreeNode
	cur := root
	for cur != nil {
		prev = cur
		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if val < prev.Val {
		prev.Left = &TreeNode{Val: val}
	} else {
		prev.Right = &TreeNode{Val: val}
	}
	return root
}
