package leetcode654

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsSameTree(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}), utils.CreateTreeByArray([]int{6, 3, 5, utils.NULL, 2, 0, utils.NULL, utils.NULL, 1})) {
		t.Error("test error")
	}
}
