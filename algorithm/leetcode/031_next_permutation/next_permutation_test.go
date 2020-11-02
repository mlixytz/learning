package leetcode031

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	l1 := []int{1, 3, 2}
	nextPermutation(l1)
	if !utils.IsArrayEqual(l1, []int{2, 1, 3}) {
		t.Error("test error")
	}
}
