package leetcode530

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if getMinimumDifference(generateBST([]int{1, 2, 3})) != 1 {
		t.Error("test error")
	}
	if getMinimumDifference(generateBST([]int{104, 227, 236, 701, 911})) != 9 {
		fmt.Println(getMinimumDifference(generateBST([]int{104, 227, 236, 701, 911})))
		t.Error("test error")
	}
}

func generateBST(s []int) *TreeNode {
	if len(s) == 1 {
		return &TreeNode{Val: s[0]}
	}
	root := len(s) / 2
	var left, right *TreeNode
	if len(s)-root == 1 {
		right = nil
	} else {
		right = generateBST(s[root:len(s)])
	}
	if root == 1 {
		left = nil
	} else {
		left = generateBST(s[0:root])
	}

	node := TreeNode{s[root], left, right}
	return &node
}
