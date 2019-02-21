package leetcode671

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil {
		return -1
	}
	left := root.Left.Val
	right := root.Right.Val

	if root.Val == root.Left.Val {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Val == root.Right.Val {
		right = findSecondMinimumValue(root.Right)
	}

	if left != -1 && right != -1 {
		if left < right {
			return left
		}
		return right
	} else if left == -1 {
		return right
	}
	return left
}
