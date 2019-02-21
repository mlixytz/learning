package leetcode137

import "testing"

func TestSolution(t *testing.T) {
	// if singleNumber([]int{2, 2, 3, 2}) != 3 {
	// 	t.Error("test error")
	// }
	// if singleNumber([]int{0, 1, 0, 1, 0, 1, 99}) != 99 {
	// 	t.Error("test error")
	// }
	if singleNumber([]int{2394, 294, 2394, 22, 2394, 294, 294}) != 22 {
		t.Error("test error")
	}
}
