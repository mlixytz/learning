package leetcode101

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if !isSymmetric(utils.CreateTreeByArray([]int{1, 2, 2, 3, 4, 4, 3})) {
		t.Error("test error")
	}
}
