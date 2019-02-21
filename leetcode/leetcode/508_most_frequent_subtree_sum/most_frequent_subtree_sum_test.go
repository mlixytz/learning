package leetcode508

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayNumerEqual(findFrequentTreeSum(utils.CreateTreeByArray([]int{5, 2, -3})), []int{2, -3, 4}) {
		t.Error("test error")
	}
	if !utils.IsArrayNumerEqual(findFrequentTreeSum(utils.CreateTreeByArray([]int{5, 2, -5})), []int{2}) {
		t.Error("test error")
	}
}
