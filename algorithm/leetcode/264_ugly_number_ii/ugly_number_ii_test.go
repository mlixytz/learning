package leetcode264

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if nthUglyNumber(10) != 12 {
		t.Error("test error")
	}
}
