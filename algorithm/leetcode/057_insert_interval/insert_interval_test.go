package leetcode057

import (
	"testing"

	"github.com/mlixytz/learning-go/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	result := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	target := [][]int{{1, 5}, {6, 9}}
	if len(result) != len(target) {
		t.Error("test erorr")
	}
	for i := 0; i < len(result); i++ {
		if !utils.IsArrayEqual(result[i], target[i]) {
			t.Error("test error")
		}
	}

	result2 := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	target2 := [][]int{{1, 2}, {3, 10}, {12, 16}}
	if len(result2) != len(target2) {
		t.Error("test erorr")
	}
	for i := 0; i < len(result2); i++ {
		if !utils.IsArrayEqual(result2[i], target2[i]) {
			t.Error("test error")
		}
	}

	result3 := insert([][]int{{1, 5}}, []int{0, 0})
	target3 := [][]int{{0, 0}, {1, 5}}
	if len(result3) != len(target3) {
		t.Error("test erorr")
	}
	for i := 0; i < len(result3); i++ {
		if !utils.IsArrayEqual(result3[i], target3[i]) {
			t.Error("test error")
		}
	}

	result4 := insert([][]int{{3, 5}, {12, 15}}, []int{6, 6})
	target4 := [][]int{{3, 5}, {6, 6}, {12, 15}}
	if len(result4) != len(target4) {
		t.Error("test erorr")
	}
	for i := 0; i < len(result4); i++ {
		if !utils.IsArrayEqual(result4[i], target4[i]) {
			t.Error("test error")
		}
	}
}
