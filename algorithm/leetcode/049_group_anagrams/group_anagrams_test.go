package leetcode049

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	ret := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	target := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	for i := 0; i < len(ret); i++ {
		if !utils.IsStringArrayEqual(target[i], ret[i]) {
			t.Error("test error")
		}
	}
}
