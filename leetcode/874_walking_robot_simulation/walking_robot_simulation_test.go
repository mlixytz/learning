package leetcode874

import "testing"

func TestSolution(t *testing.T) {
	if robotSim([]int{4, -1, 3}, [][]int{}) != 25 {
		t.Error("test error")
	}
	if robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}) != 65 {
		t.Error("test error")
	}
	obst := [][]int{{-3, 2}, {-2, 1}, {0, 1}, {-2, 4}, {-1, 0}, {-2, -3}, {0, -3}, {4, 4}, {-3, 3}, {2, 2}}
	if robotSim([]int{7, -2, -2, 7, 5}, obst) != 4 {
		t.Error("test error")
	}

}
