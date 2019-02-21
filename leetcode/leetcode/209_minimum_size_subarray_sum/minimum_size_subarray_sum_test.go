package leetcode209

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}) != 2 {
		fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
		t.Error("test error")
	}
}
