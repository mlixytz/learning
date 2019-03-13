package leetcode671

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if findSecondMinimumValue(utils.CreateTreeByArray([]int{2, 2, 5, utils.NULL, utils.NULL, 5, 7})) != 5 {
		t.Error("test error")
	}
	if findSecondMinimumValue(utils.CreateTreeByArray([]int{2, 2, 2})) != -1 {
		t.Error("test error")
	}
	if findSecondMinimumValue(utils.CreateTreeByArray([]int{2})) != -1 {
		t.Error("test error")
	}
	if findSecondMinimumValue(utils.CreateTreeByArray([]int{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1})) != 2 {
		t.Error("test error")
	}
}
