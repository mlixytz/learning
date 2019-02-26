package leetcode238

import (
	"testing"

	"github.com/mlixytz/go-algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6}) {
		t.Error("test error")
	}
}
