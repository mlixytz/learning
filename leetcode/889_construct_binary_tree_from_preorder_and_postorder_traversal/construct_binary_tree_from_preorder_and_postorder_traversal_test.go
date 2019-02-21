package leetcode889

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsSameTree(constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}), utils.CreateTreeByArray([]int{1, 2, 3, 4, 5, 6, 7})) {
		t.Error("test error")
	}
}
