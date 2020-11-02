package leetcode054

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5}) {
		t.Error("test error")
	}
	if !utils.IsArrayEqual(spiralOrder([][]int{{1, 2}, {3, 4}}), []int{1, 2, 4, 3}) {
		t.Error("test error")
	}
}
