package leetcode655

import (
	"strconv"

	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func printTree(root *TreeNode) [][]string {
	M, N := getMN(root)
	res := make([][]string, M)
	for i := 0; i < M; i++ {
		res[i] = make([]string, N)
	}
	updateOutput(root, 0, 0, N-1, res)
	return res

}

func updateOutput(root *TreeNode, row, left, right int, res [][]string) {
	if root == nil {
		return
	}
	mid := (left + right) / 2
	res[row][mid] = strconv.Itoa(root.Val)
	updateOutput(root.Left, row+1, left, mid-1, res)
	updateOutput(root.Right, row+1, mid+1, right, res)
}

func getMN(root *TreeNode) (int, int) {
	M := getHeight(root)
	N := 1
	for i := 0; i < M; i++ {
		N *= 2
	}
	return M, N - 1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := getHeight(root.Left)
	rightH := getHeight(root.Right)
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}
