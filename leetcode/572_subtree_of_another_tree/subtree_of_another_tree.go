package leetcode572

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {
	var r1, r2 bool
	if isSameTree(s, t) {
		return true
	}
	if s.Left != nil {
		r1 = isSubtree(s.Left, t)
	}
	if s.Right != nil {
		r2 = isSubtree(s.Right, t)
	}
	return r1 || r2
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if (s != nil && t != nil) && (s.Val == t.Val) {
		return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
	}
	return false
}
