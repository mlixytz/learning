package leetcode662

/*
   二叉树用数组表示，每次递归计算当前level的节点索引的宽度（索引 - 当前层的开始节点 + 1）
*/

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func widthOfBinaryTree(root *TreeNode) int {
	list := make([]int, 0)
	return dfs(root, 0, 1, &list)
}

func dfs(root *TreeNode, level, order int, list *[]int) int {
	if root == nil {
		return 0
	}
	if len(*list) == level {
		*list = append(*list, order)
	}
	cur := order - (*list)[level] + 1
	l := dfs(root.Left, level+1, 2*order, list)
	r := dfs(root.Right, level+1, 2*order+1, list)
	if l > cur {
		cur = l
	}
	if r > cur {
		cur = r
	}
	return cur
}
