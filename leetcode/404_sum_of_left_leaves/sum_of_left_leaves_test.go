package leetcode404

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if sumOfLeftLeaves(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5})) != 4 {
		t.Error("test error")
	}
	if sumOfLeftLeaves(utils.CreateTreeByArray([]int{3, 9, 20, utils.NULL, utils.NULL, 15, 7})) != 24 {
		t.Error("test error")
	}
}
