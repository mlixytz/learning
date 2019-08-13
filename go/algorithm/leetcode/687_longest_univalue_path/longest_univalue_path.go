package leetcode687

/* longestPath 记录从当前节点开始的最长同值路径（左测的或者右侧的）。全局变量res 记录当前最长的路径 */

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	longestPath(root, &res)
	return res
}

func longestPath(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	var left, right int
	if root.Left != nil {
		_left := longestPath(root.Left, res)

		if root.Left.Val == root.Val {
			left += _left + 1
		}
	}

	if root.Right != nil {
		_right := longestPath(root.Right, res)

		if root.Right.Val == root.Val {
			right += _right + 1
		}
	}

	if (left + right) > *res {
		*res = left + right
	}

	if left > right {
		return left
	}
	return right
}
