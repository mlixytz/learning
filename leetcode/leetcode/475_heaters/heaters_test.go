package leetcode475

import "testing"

func TestSolution(t *testing.T) {
	if findRadius([]int{1, 2, 3}, []int{2}) != 1 {
		t.Error("test error")
	}
	if findRadius([]int{1, 2, 3, 4}, []int{1, 4}) != 1 {
		t.Error("test error")
	}
	if findRadius([]int{1, 2, 3, 4}, []int{1}) != 3 {
		t.Error("test error")
	}
	if findRadius([]int{}, []int{}) != 0 {
		t.Error("test error")
	}
	if findRadius([]int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923},
		[]int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}) != 161834419 {
		t.Error("test error")
	}
}
