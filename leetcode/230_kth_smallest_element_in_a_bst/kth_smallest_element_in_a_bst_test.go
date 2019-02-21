package leetcode230

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if kthSmallest(utils.CreateTreeByArray([]int{3, 1, 4, utils.NULL, 2}), 1) != 1 {
		t.Error("test error")
	}
	if kthSmallest(utils.CreateTreeByArray([]int{5, 3, 6, 2, 4, utils.NULL, utils.NULL, 1}), 3) != 3 {
		t.Error("test error")
	}
}
