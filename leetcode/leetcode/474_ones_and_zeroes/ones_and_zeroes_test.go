package leetcode474

import "testing"

func TestSolution(t *testing.T) {
	if findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3) != 4 {
		t.Error("test error")
	}
	if findMaxForm([]string{"10", "0", "1"}, 1, 1) != 2 {
		t.Error("test error")
	}
}
