package leetcode606

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if tree2str(utils.CreateTreeByArray([]int{1, 2, 3, 4})) != "1(2(4))(3)" {
		t.Error("test error")
	}
	if tree2str(utils.CreateTreeByArray([]int{1, 2, 3, utils.NULL, 4})) != "1(2()(4))(3)" {
		t.Error("test error")
	}
}
