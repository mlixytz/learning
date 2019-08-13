package leetcode501

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func findMode(root *TreeNode) []int {
	var modeCount, currVal, currCount int
	ans := make([]int, 0)
	inorder(root, &ans, &modeCount, &currVal, &currCount)
	return ans
}

func inorder(root *TreeNode, ans *[]int, modeCount, currVal, currCount *int) {
	if root == nil {
		return
	}

	inorder(root.Left, ans, modeCount, currVal, currCount)

	if root.Val == *currVal {
		*currCount++
	} else {
		*currVal = root.Val
		*currCount = 1
	}

	if *currCount > *modeCount {
		*modeCount = *currCount
		*ans = nil
	}
	if *currCount == *modeCount {
		*ans = append(*ans, *currVal)
	}

	inorder(root.Right, ans, modeCount, currVal, currCount)
}
