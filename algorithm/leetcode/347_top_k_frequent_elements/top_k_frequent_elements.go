package leetcode347

func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	bucket := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		bucket[i] = make([]int, 0)
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}
	ans := make([]int, 0)
	for i := n; i >= 0; i-- {
		if len(bucket[i]) > 0 && k > len(ans) {
			ans = append(ans, bucket[i]...)
		}
	}
	return ans
}
