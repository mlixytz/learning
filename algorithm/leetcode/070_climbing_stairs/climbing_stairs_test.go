package leetcode070

import "testing"

func TestSolution(t *testing.T) {
	if climbStairs(1) != 1 {
		t.Error("test error")
	}
	if climbStairs(2) != 2 {
		t.Error("test error")
	}
	if climbStairs(3) != 3 {
		t.Error("test error")
	}
	if climbStairs(4) != 5 {
		t.Error("test error")
	}

}
