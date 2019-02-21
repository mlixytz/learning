package leetcode005

func longestPalindrome(s string) string {
	n := len(s)
	var l, h int
	if n < 2 {
		return s
	}
	expandAroundCenter := func(s string, i, j int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
		}
		if j-i > h-l {
			l, h = i, j
		}
	}
	for i := 0; i < n; i++ {
		expandAroundCenter(s, i, i)
		expandAroundCenter(s, i, i+1)
	}
	return s[l+1 : h]
}
