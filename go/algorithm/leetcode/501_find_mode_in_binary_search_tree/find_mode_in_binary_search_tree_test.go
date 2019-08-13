package leetcode501

import (
	"fmt"
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	root := utils.CreateTreeByArray([]int{1, utils.NULL, 2, 2})
	if !utils.IsArrayEqual(findMode(root), []int{2}) {
		t.Error("test error")
	}
	root2 := utils.CreateTreeByArray([]int{1, utils.NULL, 2})
	if !utils.IsArrayEqual(findMode(root2), []int{1, 2}) {
		t.Error("test error")
	}
	root3 := utils.CreateTreeByArray([]int{2, 1})
	if !utils.IsArrayEqual(findMode(root3), []int{1, 2}) {
		fmt.Println(findMode(root3))
		t.Error("test error")
	}
}
