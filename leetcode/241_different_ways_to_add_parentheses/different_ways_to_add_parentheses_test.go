package leetcode241

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if !compareTwoSlice(diffWaysToCompute("2-1-1"), []int{0, 2}) {
		fmt.Println(diffWaysToCompute("2-1-1"))
		t.Error("test error")
	}
	if !compareTwoSlice(diffWaysToCompute("2*3-4*5"), []int{-34, -14, -10, -10, 10}) {
		fmt.Println(diffWaysToCompute("2*3-4*5"))
		t.Error("test error")
	}
}

func compareTwoSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	bmap := make(map[int]bool, len(b))

	for _, val := range b {
		bmap[val] = true
	}
	for _, val := range a {
		if _, ok := bmap[val]; ok {
			continue
		}
		return false
	}
	return true
}
