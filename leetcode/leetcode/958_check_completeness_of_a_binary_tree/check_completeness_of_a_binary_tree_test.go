package leetcode958

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !isCompleteTree(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5, 6})) {
		t.Error("test error")
	}
	if isCompleteTree(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5, utils.NULL, 7})) {
		t.Error("test error")
	}
}
