package leetcode951

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	// 如果root1 不等于 root2 返回false
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	if root1.Val != root2.Val {
		return false
	}

	return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) || flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
}
