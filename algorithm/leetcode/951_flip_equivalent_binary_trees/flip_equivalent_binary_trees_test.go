package leetcode951

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !flipEquiv(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5, 6, utils.NULL, utils.NULL, utils.NULL, 7, 8}), utils.CreateTreeByArray([]int{1, 3, 2, utils.NULL, 6, 4, 5, utils.NULL, utils.NULL, utils.NULL, utils.NULL, 8, 7})) {
		t.Error("test error")
	}
}
