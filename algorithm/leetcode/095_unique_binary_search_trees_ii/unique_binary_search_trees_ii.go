package leetcode095

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return genTrees(1, n)
}

func genTrees(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	if start == end {
		return []*TreeNode{&TreeNode{Val: start}}
	}
	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := genTrees(start, i-1)
		right := genTrees(i+1, end)

		for _, l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{Val: i, Left: l, Right: r})
			}
		}
	}
	return res
}
