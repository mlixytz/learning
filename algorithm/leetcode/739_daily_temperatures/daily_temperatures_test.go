package leetcode739

import (
	"testing"

	"github.com/mlixytz/learning-go/algorithm/leetcode/utils"
)

func TestDailyTemperatures(t *testing.T) {
	if !utils.IsArrayEqual(dailyTemperatures([]int{}), []int{}) {
		t.Error("test error")
	}
	if !utils.IsArrayEqual(dailyTemperatures([]int{1}), []int{0}) {
		t.Error("test error")
	}
	if !utils.IsArrayEqual(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0}) {
		t.Error("test error")
	}
}
