package leetcode005

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if longestPalindrome("babad") != "bab" && longestPalindrome("babad") != "aba" {
		t.Error("test error")
	}
	if longestPalindrome("cbbd") != "bb" {
		fmt.Println(longestPalindrome("cbbd"))
		t.Error("test error")
	}
	if longestPalindrome("ccc") != "ccc" {
		t.Error("test error")
	}
	if longestPalindrome("bananas") != "anana" {
		t.Error("test error")
	}
	if longestPalindrome("caba") != "aba" {
		t.Error("test error")
	}
	if longestPalindrome("aaabaaaa") != "aaabaaa" {
		t.Error("test error")
	}
	if longestPalindrome("ababababababa") != "ababababababa" {
		t.Error("test error")
	}
	if longestPalindrome("bb") != "bb" {
		t.Error("test error")
	}
}
