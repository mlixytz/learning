package utils

import (
	"strconv"
)

// TreeNode 是一个树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 用来标示空节点的值
var NULL = -1 << 63

// CreateTreeByArray 通过给定的数组来创建树
func CreateTreeByArray(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	i := 1
	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// 左节点
		if nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		// 右节点
		if i < len(nums) && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// ToString 将二叉树转成字符串
func (root *TreeNode) ToString() string {
	if root == nil {
		return ""
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	var str string

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			str += "NULL, "
			continue
		}
		str += strconv.Itoa(node.Val) + ", "

		// 添加左，右节点
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	return str[:len(str)-2]
}

// IsSameTree 判断两棵树是否相同
func IsSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if (s != nil && t != nil) && (s.Val == t.Val) {
		return IsSameTree(s.Left, t.Left) && IsSameTree(s.Right, t.Right)
	}
	return false
}
