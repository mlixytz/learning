package leetcode712

import "testing"

func TestSolution(t *testing.T) {
	if minimumDeleteSum("sea", "eat") != 231 {
		t.Error("test error")
	}
	if minimumDeleteSum("delete", "leet") != 403 {
		t.Error("test error")
	}
}
