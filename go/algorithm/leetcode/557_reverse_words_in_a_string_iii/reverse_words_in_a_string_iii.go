package leetcode557

func reverseWords(s string) string {
	ans := []byte(s)
	for i, j := 0, 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			k := i - 1
			for j < k {
				ans[j], ans[k] = ans[k], ans[j]
				j++
				k--
			}
			j = i + 1
		}
	}
	return string(ans)
}
