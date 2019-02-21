package leetcode309

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if maxProfit([]int{1, 2, 3, 0, 2}) != 3 {
		fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
		t.Error("test error")
	}
}
