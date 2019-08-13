package leetcode538

/* 先加右子树，再加根结点，再加左子树 因为（右>根>左) */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convert(root, &sum)
	return root
}

func convert(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	convert(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	convert(root.Left, sum)
}
