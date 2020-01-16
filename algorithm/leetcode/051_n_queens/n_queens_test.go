package leetcode051

import (
	"testing"

	"github.com/mlixytz/learning-go/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	res := solveNQueens(4)
	target := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}
	for i := 0; i < 2; i++ {
		if !utils.IsStringArrayEqual(res[i], target[i]) {
			t.Error("test error")
		}
	}
}
