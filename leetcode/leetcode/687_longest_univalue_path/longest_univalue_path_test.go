package leetcode687

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if longestUnivaluePath(utils.CreateTreeByArray([]int{5, 4, 5, 1, 1, 5})) != 2 {
		t.Error("test error")
	}
	if longestUnivaluePath(utils.CreateTreeByArray([]int{1, 4, 5, 4, 4, 5})) != 2 {
		t.Error("test error")
	}
}
