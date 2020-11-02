package leetcode101

import "github.com/mlixytz/learning/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorTree(root.Left, root.Right)
}

func isMirrorTree(t1, t2 *TreeNode) bool {
	if t1 == nil {
		return t2 == nil
	}
	if t2 == nil {
		return t1 == nil
	}
	if t1.Val != t2.Val {
		return false
	}
	return isMirrorTree(t1.Left, t2.Right) && isMirrorTree(t1.Right, t2.Left)
}
