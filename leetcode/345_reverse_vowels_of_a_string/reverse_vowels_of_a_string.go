package leetcode345

func reverseVowels(s string) string {
	ans := []byte(s)
	l := 0
	r := len(ans) - 1
	for l < r {
		for l < r && s[l] != 'a' && s[l] != 'e' && s[l] != 'i' && s[l] != 'o' && s[l] != 'u' && s[l] != 'A' && s[l] != 'E' && s[l] != 'I' && s[l] != 'O' && s[l] != 'U' {
			l++
		}
		for l < r && s[r] != 'a' && s[r] != 'e' && s[r] != 'i' && s[r] != 'o' && s[r] != 'u' && s[r] != 'A' && s[r] != 'E' && s[r] != 'I' && s[r] != 'O' && s[r] != 'U' {
			r--
		}
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return string(ans)
}
