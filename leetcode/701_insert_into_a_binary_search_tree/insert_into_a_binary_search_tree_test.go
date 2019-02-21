package leetcode701

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsSameTree(insertIntoBST(utils.CreateTreeByArray([]int{4, 2, 7, 1, 3}), 5), utils.CreateTreeByArray([]int{4, 2, 7, 1, 3, 5})) {
		t.Error("test error")
	}
}
