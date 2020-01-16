package leetcode977

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(sortedSquares([]int{-4, -1, 0, 3, 10}), []int{0, 1, 9, 16, 100}) {
		t.Error("test error")
	}
}
