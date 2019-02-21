package leetcode322

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if coinChange([]int{1, 2, 5}, 11) != 3 {
		t.Error("test error")
	}
	if coinChange([]int{2}, 3) != -1 {
		t.Error("test error")
	}
	if coinChange([]int{1}, 0) != 0 {
		t.Error("test error")
	}
	if coinChange([]int{186, 419, 83, 408}, 6249) != 20 {
		fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
		t.Error("test error")
	}
}
