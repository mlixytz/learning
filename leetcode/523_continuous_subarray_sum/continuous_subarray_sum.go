package leetcode532

/* (a + n*k) % k == a % k */

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	index := make(map[int]int)
	index[0] = -1

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}

		idx, ok := index[sum]
		if ok {
			if i-idx > 1 {
				return true
			}
		} else {
			index[sum] = i
		}
	}
	return false
}
