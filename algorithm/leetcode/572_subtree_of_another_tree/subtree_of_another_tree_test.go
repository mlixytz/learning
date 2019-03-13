package leetcode572

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !isSubtree(utils.CreateTreeByArray([]int{3, 4, 5, 1, 2}), utils.CreateTreeByArray([]int{4, 1, 2})) {
		t.Error("test error")
	}
	if isSubtree(utils.CreateTreeByArray([]int{3, 4, 5, 1, 2, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 0}), utils.CreateTreeByArray([]int{4, 1, 2})) {
		t.Error("test error")
	}
}
