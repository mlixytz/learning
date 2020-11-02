package leetcode437

import "github.com/mlixytz/learning/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSumStartRoot(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumStartRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var count int
	if root.Val == sum {
		count++
	}
	count += pathSumStartRoot(root.Left, sum-root.Val)
	count += pathSumStartRoot(root.Right, sum-root.Val)
	return count
}
