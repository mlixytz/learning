package leetcode230

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	count, num := k, 0
	inOrder(root, &count, &num)
	return num
}

func inOrder(root *TreeNode, count, num *int) {
	if root.Left != nil {
		inOrder(root.Left, count, num)
	}
	*count--
	if *count == 0 {
		*num = root.Val
		return
	}
	if root.Right != nil {
		inOrder(root.Right, count, num)
	}
}
