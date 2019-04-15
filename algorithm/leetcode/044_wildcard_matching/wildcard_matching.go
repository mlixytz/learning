package leetcode044

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	match := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		match[i] = make([]bool, n+1)
	}
	match[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			match[0][j] = true
		} else {
			break
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == s[i] || p[j] == '?' {
				match[i+1][j+1] = match[i][j]
			} else if p[j] == '*' {
				match[i+1][j+1] = match[i][j+1] || match[i+1][j] // * 不为空 || * 为空
			}
		}
	}
	return match[m][n]
}
