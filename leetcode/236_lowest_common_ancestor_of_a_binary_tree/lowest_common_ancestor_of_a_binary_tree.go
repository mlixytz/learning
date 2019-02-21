package leetcode236

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil && right == nil {
		return nil
	} else {
		if left == nil {
			return right
		}
		return left
	}
}
