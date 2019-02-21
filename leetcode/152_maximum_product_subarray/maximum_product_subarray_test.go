package leetcode152

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if maxProduct([]int{2, 3, -2, 4}) != 6 {
		t.Error("test error")
	}
	if maxProduct([]int{-2, 0, -1}) != 0 {
		t.Error("test error")
	}
	if maxProduct([]int{0, 2}) != 2 {
		fmt.Println(maxProduct([]int{0, 2}))
		t.Error("test error")
	}
	if maxProduct([]int{-4, -3}) != 12 {
		t.Error("test error")
	}
	if maxProductBS([]int{2, 3, -2, 4}) != 6 {
		t.Error("test error")
	}
	if maxProductBS([]int{-2, 0, -1}) != 0 {
		t.Error("test error")
	}
	if maxProductBS([]int{0, 2}) != 2 {
		t.Error("test error")
	}
	if maxProductBS([]int{-4, -3}) != 12 {
		t.Error("test error")
	}
}
