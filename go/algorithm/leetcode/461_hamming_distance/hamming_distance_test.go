package leetcode461

import "testing"

func TestSolution(t *testing.T) {
	if hammingDistance(1, 4) != 2 {
		t.Error("test error")
	}
	if hammingDistance(3, 1) != 1 {
		t.Error("test error")
	}
}
