package leetcode110

import (
	"testing"

	"github.com/mlixytz/learnging/go-algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if isBalanced(utils.CreateTreeByArray([]int{1, 2, 2, 3, 3, utils.Null, utils.Null, 4, 4})) {
		t.Error("test error")
	}
}
