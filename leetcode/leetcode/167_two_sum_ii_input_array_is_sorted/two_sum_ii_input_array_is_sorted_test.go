package leetcode167

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2}) {
		t.Error("test error")
	}
}
