package leetcode030

import (
	"testing"

	"github.com/mlixytz/learning-go/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}), []int{}) {
		t.Error("test error")
	}
	if !utils.IsArrayEqual(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}), []int{8}) {
		t.Error("test error")
	}
}
