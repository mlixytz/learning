package leetcode264

/*
   解发同超级丑数
*/

func nthUglyNumber(n int) int {
	var t2, t3, t5 int
	ugly := make([]int, n)
	ugly[0] = 1
	for i := 1; i < n; i++ {
		m := min(2*ugly[t2], min(3*ugly[t3], 5*ugly[t5]))
		if m == 2*ugly[t2] {
			t2++
		}
		if m == 3*ugly[t3] {
			t3++
		}
		if m == 5*ugly[t5] {
			t5++
		}
		ugly[i] = m
	}
	return ugly[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
