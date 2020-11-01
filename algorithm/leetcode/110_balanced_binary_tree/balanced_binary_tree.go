package leetcode110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if maxDepthOfTree(root) == -1 {
		return false
	}
	return true
}

func maxDepthOfTree(head *TreeNode) int {
	if head == nil {
		return 0
	}
	left := maxDepthOfTree(head.Left)
	right := maxDepthOfTree(head.Right)

	if left == -1 || right == -1 || abs(left, right) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return 0 - res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
