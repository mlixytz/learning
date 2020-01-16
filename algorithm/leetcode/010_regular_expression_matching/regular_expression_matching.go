package leetcode010

/*

1, If p.charAt(j) == s.charAt(i) :  dp[i][j] = dp[i-1][j-1];
2, If p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1];
3, If p.charAt(j) == '*':
   here are two sub conditions:
               1   if p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2]  //in this case, a* only counts as empty
               2   if p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.':
                              dp[i][j] = dp[i-1][j]    //in this case, a* counts as multiple a
                           or dp[i][j] = dp[i][j-1]   // in this case, a* counts as single a
			   or dp[i][j] = dp[i][j-2]   // in this case, a* counts as empty

*/

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return false
	}

	n, m := len(p), len(s)

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := 0; j < n; j++ {
		if p[j] == '*' {
			if dp[0][j] || (j > 0 && dp[0][j-1]) {
				dp[0][j+1] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == s[i] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1] // a* 为空
				} else {
					dp[i+1][j+1] = (dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]) // a* 多字符 || * 单子符 || a* 为空
				}
			}
		}
	}
	return dp[m][n]
}
