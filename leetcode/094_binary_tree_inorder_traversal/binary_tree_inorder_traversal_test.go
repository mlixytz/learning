package leetcode094

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(inorderTraversal(utils.CreateTreeByArray([]int{1, utils.NULL, 2, 3})), []int{1, 3, 2}) {
		t.Error("test erorr")
	}
}
