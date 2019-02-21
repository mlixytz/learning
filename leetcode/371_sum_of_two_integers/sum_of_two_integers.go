package leetcode371

func getSum(a, b int) int {
	if b == 0 {
		return a
	}
	return getSum(a^b, (a&b)<<1)
}
