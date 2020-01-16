package leetcode089

import (
	"testing"

	"github.com/mlixytz/go-algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(grayCode(1), []int{0, 1}) {
		t.Error("test error")
	}
	if !utils.IsArrayEqual(grayCode(2), []int{0, 1, 3, 2}) {
		t.Error("test error")
	}
}
