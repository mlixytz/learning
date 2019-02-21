package leetcode746

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	f1, f2 := 0, 0
	for i := 0; i < size; i++ {
		f1, f2 = f2, cost[i]+min(f1, f2)
	}
	return min(f1, f2)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
