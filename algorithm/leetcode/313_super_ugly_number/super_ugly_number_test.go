package leetcode313

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if nthSuperUglyNumber(12, []int{2, 7, 13, 19}) != 32 {
		t.Error("test error")
	}
}
