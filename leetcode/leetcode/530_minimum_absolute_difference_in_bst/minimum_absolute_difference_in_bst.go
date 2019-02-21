package leetcode530

import (
	"math"
)

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type state struct {
	md, prev int
}

func getMinimumDifference(root *TreeNode) int {
	st := state{math.MaxInt32, -1}
	minidiff(root, &st)
	return st.md
}

func minidiff(root *TreeNode, st *state) {
	if root == nil {
		return
	}

	minidiff(root.Left, st)

	if st.prev >= 0 {
		st.md = min(st.md, root.Val-st.prev)
	}
	st.prev = root.Val

	minidiff(root.Right, st)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
