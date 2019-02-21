package leetcode688

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if knightProbability(3, 2, 0, 0) != 0.0625 {
		t.Error("test error")
	}
}
