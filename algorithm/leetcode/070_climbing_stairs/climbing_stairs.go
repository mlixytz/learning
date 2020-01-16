package leetcode070

func climbStairs(n int) int {
	sum := make([]int, n)
	if n < 2 {
		return 1
	}
	sum[0] = 1
	sum[1] = 2

	for i := 2; i < n; i++ {
		sum[i] = sum[i-1] + sum[i-2]
	}
	return sum[n-1]
}
