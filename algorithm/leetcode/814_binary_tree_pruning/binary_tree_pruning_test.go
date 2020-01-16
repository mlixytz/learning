package leetcode814

import (
	"fmt"
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsSameTree(pruneTree(utils.CreateTreeByArray([]int{1, utils.NULL, 0, 0, 1})), utils.CreateTreeByArray([]int{1, utils.NULL, 0, utils.NULL, 1})) {
		fmt.Println(pruneTree(utils.CreateTreeByArray([]int{1, utils.NULL, 0, 0, 1})).ToString())
		t.Error("test error")
	}
	if !utils.IsSameTree(pruneTree(utils.CreateTreeByArray([]int{1, 0, 1, 0, 0, 0, 1})), utils.CreateTreeByArray([]int{1, utils.NULL, 1, utils.NULL, 1})) {
		t.Error("test error")
	}
	if !utils.IsSameTree(pruneTree(utils.CreateTreeByArray([]int{1, 1, 0, 1, 1, 0, 1, 0})), utils.CreateTreeByArray([]int{1, 1, 0, 1, 1, utils.NULL, 1})) {
		t.Error("test error")
	}
}
