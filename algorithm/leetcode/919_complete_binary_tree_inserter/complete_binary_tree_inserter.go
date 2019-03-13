package leetcode919

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

type CBTInserter struct {
	array []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	array := make([]*TreeNode, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 && queue[0] != nil {
		cur := queue[0]
		array = append(array, cur)
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
		queue = queue[1:]
	}
	return CBTInserter{array}
}

func (this *CBTInserter) Insert(v int) int {
	size := len(this.array)
	node := &TreeNode{Val: v}
	var rindex int
	if size&1 != 0 { // 奇数， 左节点 2 x i + 1
		rindex = (size - 1) / 2
		this.array[rindex].Left = node
	} else { // 偶数，右节点 2 x i + 2
		rindex = (size - 2) / 2
		this.array[rindex].Right = node
	}
	this.array = append(this.array, node)
	return this.array[rindex].Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	if len(this.array) > 0 {
		return this.array[0]
	}
	return nil
}
