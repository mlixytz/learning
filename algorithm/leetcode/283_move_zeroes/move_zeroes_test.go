package leetcode283

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	if !utils.IsArrayEqual(nums, []int{1, 3, 12, 0, 0}) {
		t.Error("test error")
	}
}
