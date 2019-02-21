package leetcode543

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	maxDepth(root, &max)
	return max
}

func maxDepth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left, max)
	right := maxDepth(root.Right, max)
	if left+right > *max {
		*max = left + right
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
