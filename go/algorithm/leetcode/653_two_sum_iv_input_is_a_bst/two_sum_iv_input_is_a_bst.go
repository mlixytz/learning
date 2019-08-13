package leetcode653

/* 将bst树转换为有序数组，用双指针找到目标值 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	nums := make([]int, 0)
	inorder(root, &nums)
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] == k {
			return true
		}
		if nums[i]+nums[j] < k {
			i++
		} else {
			j--
		}
	}
	return false
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}
