package leetcode222

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
