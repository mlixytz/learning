package leetcode367

func isPerfectSquare(num int) bool {
	l, h := 0, num
	for l <= h {
		m := l + (h-l)/2
		if m*m > num {
			h = m - 1
		} else if m*m < num {
			l = m + 1
		} else {
			return true
		}
	}
	return false
}
