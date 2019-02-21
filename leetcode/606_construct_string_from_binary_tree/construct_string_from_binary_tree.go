package leetcode606

import (
	"strconv"

	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	ans := strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		return ans
	}
	ans += "(" + tree2str(t.Left) + ")"
	if t.Right != nil {
		ans += "(" + tree2str(t.Right) + ")"
	}
	return ans
}
