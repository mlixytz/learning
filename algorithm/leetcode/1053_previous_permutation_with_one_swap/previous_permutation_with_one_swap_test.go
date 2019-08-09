package leetcode1053

import (
	"testing"

	"github.com/mlixytz/learning-go/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(prevPermOpt1([]int{1, 9, 4, 6, 7}), []int{1, 7, 4, 6, 9}) {
		t.Error("test error")
	}
	if !utils.IsArrayEqual(prevPermOpt1([]int{3, 1, 1, 3}), []int{1, 3, 1, 3}) {
		t.Error("test error")
	}
}
