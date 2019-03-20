package leetcode006

func convert(s string, numRows int) string {
	n := len(s)
	if n < 2 {
		return s
	}
	rets := make([][]byte, numRows)

	i := 0
	for i < n {
		for idx := 0; idx < numRows && i < n; idx++ { // 向下
			rets[idx] = append(rets[idx], s[i])
			i++
		}
		for idx := numRows - 2; idx >= 1 && i < n; idx-- { // 向上
			rets[idx] = append(rets[idx], s[i])
			i++
		}
	}

	ret := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		ret = append(ret, rets[i]...)
	}
	return string(ret)
}
