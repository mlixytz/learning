package leetcode662

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if widthOfBinaryTree(utils.CreateTreeByArray([]int{1, 3, 2, 5, 3, utils.NULL, 9})) != 4 {
		t.Error("test error")
	}
	if widthOfBinaryTree(utils.CreateTreeByArray([]int{1, 3, utils.NULL, 5, 3})) != 2 {
		t.Error("test error")
	}
}
