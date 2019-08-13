package leetcode958

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func isCompleteTree(root *TreeNode) bool {
	bfs := make([]*TreeNode, 0)
	bfs = append(bfs, root)
	for bfs[0] != nil {
		cur := bfs[0]
		bfs = append(bfs, cur.Left)
		bfs = append(bfs, cur.Right)
		bfs = bfs[1:]
	}
	for len(bfs) > 0 && bfs[0] == nil {
		bfs = bfs[1:]
	}
	return len(bfs) == 0
}
