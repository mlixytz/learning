package leetcode437

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if pathSum(utils.CreateTreeByArray([]int{10, 5, -3, 3, 2, utils.NULL, 11, 3, -2, utils.NULL, 1}), 8) != 3 {
		t.Error("test error")
	}
}
