package leetcode018

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	ret := [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}
	for i, val := range fourSum([]int{1, 0, -1, 0, -2, 2}, 0) {
		if !utils.IsArrayEqual(val, ret[i]) {
			t.Error("test error")
		}
	}
}
