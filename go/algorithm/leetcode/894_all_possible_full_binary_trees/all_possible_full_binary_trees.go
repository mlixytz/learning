package leetcode894

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func allPossibleFBT(N int) []*TreeNode {
	res := []*TreeNode{}

	if N == 1 {
		res = append(res, &TreeNode{Val: 0})
		return res
	}
	for i := 1; i < N-1; i += 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(N - 1 - i)
		for _, v1 := range left {
			for _, v2 := range right {
				res = append(res, &TreeNode{Val: 0, Left: v1, Right: v2})
			}
		}
	}
	return res
}
