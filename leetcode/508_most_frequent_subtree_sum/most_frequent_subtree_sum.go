package leetcode508

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func findFrequentTreeSum(root *TreeNode) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	max := 0
	postOrder(root, &m, &max)
	for k, v := range m {
		if max == v {
			res = append(res, k)
		}
	}
	return res
}

func postOrder(root *TreeNode, m *map[int]int, max *int) int {
	if root == nil {
		return 0
	}
	left := postOrder(root.Left, m, max)
	right := postOrder(root.Right, m, max)
	sum := left + right + root.Val
	(*m)[sum]++
	if (*m)[sum] > *max {
		*max = (*m)[sum]
	}
	return sum
}
