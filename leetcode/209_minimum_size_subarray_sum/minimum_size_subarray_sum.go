package leetcode209

func minSubArrayLen(s int, nums []int) int {
	i, j, ans, tmps := 0, 0, len(nums)+1, 0
	for j < len(nums) {
		tmps += nums[j]
		j++
		for tmps >= s {
			if j-i < ans {
				ans = j - i
			}
			tmps -= nums[i]
			i++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
