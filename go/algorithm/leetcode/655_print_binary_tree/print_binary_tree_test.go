package leetcode655

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	res := printTree(utils.CreateTreeByArray([]int{1, 2}))

	op := [][]string{{"", "1", ""}, {"2", "", ""}}

	if len(res) != len(op) {
		t.Error("test error")
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] != op[i][j] {
				t.Error("test error")
				return
			}
		}
	}
}
