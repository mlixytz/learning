package leetcode652

import (
	"strconv"

	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string]int)
	res := make([]*TreeNode, 0)
	preOrder(root, m, &res)
	return res
}

func preOrder(root *TreeNode, m map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	serial := strconv.Itoa(root.Val) + "," + preOrder(root.Left, m, res) + "," + preOrder(root.Right, m, res)
	if val, ok := m[serial]; ok && val == 1 {
		*res = append(*res, root)
	}
	m[serial]++
	return serial
}
