package leetcode048

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	target := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	for i := range matrix {
		if !utils.IsArrayNumerEqual(matrix[i], target[i]) {
			t.Error("test error")
		}
	}

}
