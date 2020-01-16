package leetcode017

import (
	"testing"

	"github.com/mlixytz/learning-go/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsStringArrayEqual(letterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}) {
		t.Error("test error")
	}
}
