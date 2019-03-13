package leetcode441

import "testing"

func TestSolution(t *testing.T) {
	if arrangeCoins(0) != 0 {
		t.Error("test error")
	}
	if arrangeCoins(1) != 1 {
		t.Error("test error")
	}
	t.Log(arrangeCoins(20))
	if arrangeCoins(20) != 5 {
		t.Error("test error")
	}
	if arrangeCoins(21) != 6 {
		t.Error("test error")
	}
	if arrangeCoins(22) != 6 {
		t.Error("test error")
	}
}
