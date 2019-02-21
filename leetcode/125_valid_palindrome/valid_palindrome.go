package leetcode125

import "strings"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	s = strings.ToLower(s)

	for l < r {
		if !isChar(s[l]) {
			l++
		} else if !isChar(s[r]) {
			r--
		} else {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
