package leetcode124

import (
	"testing"

	"github.com/mlixytz/go-algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if maxPathSum(utils.CreateTreeByArray([]int{1, 2, 3})) != 6 {
		t.Error("test error")
	}
	if maxPathSum(utils.CreateTreeByArray([]int{-10, 9, 20, utils.NULL, utils.NULL, 15, 7})) != 42 {
		t.Error("test error")
	}
}
