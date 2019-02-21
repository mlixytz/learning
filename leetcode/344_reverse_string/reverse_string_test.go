package leetcode344

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	if !utils.IsBytesEqual(s, []byte{'o', 'l', 'l', 'e', 'h'}) {
		t.Error("test error")
	}
}
