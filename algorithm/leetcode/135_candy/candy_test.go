package leetcode135

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := []int{1, 0, 2}
	if candy(input) != 5 {
		t.Error("test error")
	}
}
