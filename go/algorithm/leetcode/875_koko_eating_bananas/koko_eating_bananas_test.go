package leetcode875

import "testing"

func TestSolution(t *testing.T) {
	if minEatingSpeed([]int{3, 6, 7, 11}, 8) != 4 {
		t.Error("test Error")
	}
	if minEatingSpeed([]int{30, 11, 23, 4, 20}, 5) != 30 {
		t.Error("test Error")
	}
	if minEatingSpeed([]int{30, 11, 23, 4, 20}, 6) != 23 {
		t.Error("test Error")
	}
	if minEatingSpeed([]int{312884470}, 968709470) != 1 {
		t.Error("test Error")
	}
}
