package leetcode191

import "testing"

func TestSolution(t *testing.T) {
	if hammingWeight(5) != 2 {
		t.Error("test error")
	}
	if hammingWeight(17) != 2 {
		t.Error("test error")
	}
}
