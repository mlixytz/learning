package leetcode404

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	travel(root, &sum, false)
	return sum
}

func travel(root *TreeNode, sum *int, flag bool) {
	if root == nil {
		return
	}
	if root.Left != nil {
		travel(root.Left, sum, true)
	} else if flag && root.Right == nil {
		*sum += root.Val
	}
	travel(root.Right, sum, false)
}
