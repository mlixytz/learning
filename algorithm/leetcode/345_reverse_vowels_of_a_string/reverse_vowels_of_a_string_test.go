package leetcode345

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if reverseVowels("hello") != "holle" {
		t.Error("test error")
	}
}
