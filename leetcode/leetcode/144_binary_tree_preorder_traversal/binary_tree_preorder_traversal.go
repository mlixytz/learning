package leetcode144

import (
	"container/list"

	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	var cur *TreeNode
	for stack.Len() != 0 {
		cur = stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack.PushBack(cur.Right)
		}
		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}
	}
	return res
}
