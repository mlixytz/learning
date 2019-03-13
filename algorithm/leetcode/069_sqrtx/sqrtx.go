package leetcode069

func mySqrt(x int) int {
	l, h := 0, x

	for l <= h {
		m := l + (h-l)/2
		if m*m > x {
			h = m - 1
		} else if m*m < x {
			l = m + 1
		} else {
			return m
		}
	}
	return h
}
