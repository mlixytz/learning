package leetcode700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return root
	}
}
