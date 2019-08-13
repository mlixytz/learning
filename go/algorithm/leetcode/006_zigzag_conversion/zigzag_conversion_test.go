package leetcode006

import "testing"

func TestSolution(t *testing.T) {
	if convert("LEETCODEISHIRING", 4) != "LDREOEIIECIHNTSG" {
		t.Error("test error")
	}
}
