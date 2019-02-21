package leetcode450

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsSameTree(deleteNode(utils.CreateTreeByArray([]int{5, 3, 6, 2, 4, utils.NULL, 7}), 3), utils.CreateTreeByArray([]int{5, 4, 6, 2, utils.NULL, utils.NULL, 7})) {
		t.Error("test error")
	}
}
