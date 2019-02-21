package leetcode222

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if countNodes(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5, 6})) != 6 {
		t.Error("test error")
	}
}
