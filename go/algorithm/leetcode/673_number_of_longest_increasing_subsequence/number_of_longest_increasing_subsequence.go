package leetcode673

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	Len := make([]int, n)
	count := make([]int, n)
	maxLen := 1
	for i := 0; i < n; i++ {
		Len[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if Len[j]+1 > Len[i] {
					Len[i] = Len[j] + 1
					count[i] = count[j]
					if Len[i] > maxLen {
						maxLen = Len[i]
					}
				} else if Len[j]+1 == Len[i] {
					count[i] += count[j]
				}
			}
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if Len[i] == maxLen {
			sum += count[i]
		}
	}
	return sum
}
