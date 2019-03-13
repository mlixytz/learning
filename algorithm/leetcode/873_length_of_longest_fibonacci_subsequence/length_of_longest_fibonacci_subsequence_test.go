package leetcode873

import (
	"testing"
)

func TestSolution(t *testing.T) {
	// if lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}) != 5 {
	// 	t.Error("test error")
	// }
	// if lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}) != 3 {
	// 	t.Error("test error")
	// }
	// if lenLongestFibSubseq([]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}) != 5 {
	// 	t.Error("test error")
	// }
	// if lenLongestFibSubseqBS([]int{1, 2, 3, 4, 5, 6, 7, 8}) != 5 {
	// 	t.Error("test error")
	// }
	// if lenLongestFibSubseqBS([]int{1, 3, 7, 11, 12, 14, 18}) != 3 {
	// 	t.Error("test error")
	// }
	if lenLongestFibSubseqBS([]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}) != 5 {
		t.Error("test error")
	}
}
