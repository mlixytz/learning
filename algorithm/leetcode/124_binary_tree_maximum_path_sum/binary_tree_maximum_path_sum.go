package leetcode124

import (
	"github.com/mlixytz/go-algorithm/leetcode/utils"
)

type TreeNode = utils.TreeNode

func maxPathSum(root *TreeNode) int {
	sum := -1 << 31
	maxPathDown(root, &sum)
	return sum
}

func maxPathDown(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	left := max(0, maxPathDown(root.Left, sum))
	right := max(0, maxPathDown(root.Right, sum))
	*sum = max(*sum, left+right+root.Val)
	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
