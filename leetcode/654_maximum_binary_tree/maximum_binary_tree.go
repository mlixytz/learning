package leetcode654

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := findLargestNumIndex(nums)
	return &TreeNode{nums[index], constructMaximumBinaryTree(nums[:index]), constructMaximumBinaryTree(nums[index+1:])}
}

func findLargestNumIndex(nums []int) int {
	max := -(1 << 31)
	index := 0
	for i, val := range nums {
		if max < val {
			index = i
			max = val
		}
	}
	return index
}
