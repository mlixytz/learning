package leetcode022

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	res := make([]string, 0)
	backtrack(&res, n-1, n, "(")
	return res
}

func backtrack(res *[]string, pcount, ncount int, s string) {
	if pcount == 0 && ncount == 0 {
		*res = append(*res, s)
		return
	}
	if pcount > 0 {
		backtrack(res, pcount-1, ncount, s+"(")
	}
	// 如果关闭括号数量大于打开括号数量
	if ncount > pcount {
		backtrack(res, pcount, ncount-1, s+")")
	}
}
