package leetcode919

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	obj := Constructor(utils.CreateTreeByArray([]int{1}))
	if obj.Insert(2) != 1 {
		t.Error("test error")
	}
	if !utils.IsSameTree(obj.Get_root(), utils.CreateTreeByArray([]int{1, 2})) {
		t.Error("test error")
	}

	obj1 := Constructor(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5, 6}))
	if obj1.Insert(7) != 3 {
		t.Error("test error")
	}
	if obj1.Insert(8) != 4 {
		t.Error("test error")
	}
	if !utils.IsSameTree(obj1.Get_root(), utils.CreateTreeByArray([]int{1, 2, 3, 4, 5, 6, 7, 8})) {
		t.Error("test error")
	}

}
