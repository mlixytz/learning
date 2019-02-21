package leetcode231

import "testing"

func TestSolution(t *testing.T) {
	if !isPowerOfTwo(2) {
		t.Error("test error")
	}
	if !isPowerOfTwo(16) {
		t.Error("test error")
	}
	if isPowerOfTwo(218) {
		t.Error("test error")
	}
}
