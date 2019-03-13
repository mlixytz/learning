package leetcode623

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsSameTree(addOneRow(utils.CreateTreeByArray([]int{4, 2, 6, 3, 1, 5}), 1, 2), utils.CreateTreeByArray([]int{4, 1, 1, 2, utils.NULL, utils.NULL, 6, 3, 1, 5})) {
		t.Error("test error")
	}
	if !utils.IsSameTree(addOneRow(utils.CreateTreeByArray([]int{4, 2, utils.NULL, 3, 1}), 1, 3), utils.CreateTreeByArray([]int{4, 2, utils.NULL, 1, 1, 3, utils.NULL, utils.NULL, 1})) {
		t.Error("test error")
	}
}
