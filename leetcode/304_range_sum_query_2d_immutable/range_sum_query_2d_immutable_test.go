package leetcode304

import "testing"

func TestSolution(t *testing.T) {
	// obj := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	// if obj.SumRegion(2, 1, 4, 3) != 8 {
	// 	t.Error("test error")
	// }
	// if obj.SumRegion(1, 1, 2, 2) != 11 {
	// 	t.Error("test error")
	// }
	// if obj.SumRegion(1, 2, 2, 4) != 12 {
	// 	t.Error("test error")
	// }
	obj2 := Constructor([][]int{{}})
	if obj2.SumRegion(2, 1, 4, 3) != 0 {
		t.Error("test error")
	}
}
