package leetcode063

import "testing"

func TestSolution(t *testing.T) {
	obstacles := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	if uniquePathsWithObstacles(obstacles) != 2 {
		t.Error("test error")
	}
}
