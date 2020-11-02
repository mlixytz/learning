package leetcode014

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if maxDepth(utils.CreateTreeByArray([]int{3, 9, 20, utils.NULL, utils.NULL, 15, 7})) != 3 {
		t.Error("test error")
	}
}
