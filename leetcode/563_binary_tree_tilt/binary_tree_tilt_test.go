package leetcode563

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if findTilt(utils.CreateTreeByArray([]int{1, 2, 3})) != 1 {
		t.Error("test error")
	}
}
