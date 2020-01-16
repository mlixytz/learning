package leetcode563

/* 后序遍历 */

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func findTilt(root *TreeNode) int {
	sum := 0
	postOrder(root, &sum)
	return sum
}

func postOrder(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	left := postOrder(root.Left, sum)
	right := postOrder(root.Right, sum)

	res := left - right
	if res < 0 {
		*sum -= res
	} else {
		*sum += res
	}

	return left + right + root.Val
}
