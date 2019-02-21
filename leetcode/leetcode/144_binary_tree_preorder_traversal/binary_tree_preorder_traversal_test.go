package leetcode144

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(preorderTraversal(utils.CreateTreeByArray([]int{1, utils.NULL, 2, 3})), []int{1, 2, 3}) {
		t.Error("test error")
	}
}
