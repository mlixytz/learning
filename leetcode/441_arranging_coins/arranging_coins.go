package leetcode441

func arrangeCoins(n int) int {
	l, h := 0, n

	for l <= h {
		m := l + (h-l)/2
		flag := m * (m + 1) / 2
		if flag > n {
			h = m - 1
		} else if flag < n {
			l = m + 1
		} else {
			return m
		}
	}
	return h
}
