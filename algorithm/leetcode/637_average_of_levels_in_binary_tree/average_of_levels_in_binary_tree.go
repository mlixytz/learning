package leetcode637

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	queue := list.New()

	queue.PushBack(root)
	for queue.Len() != 0 {
		n := queue.Len()
		sum := 0.0
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			sum += float64(node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, (sum / float64(n)))
	}
	return result
}
