package leetcode532

func findPairs(nums []int, k int) int {
	if len(nums) == 0 || k < 0 {
		return 0
	}
	ans := 0
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	if k > 0 {
		for key := range m {
			if _, ok := m[key+k]; ok {
				ans++
			}
		}
	} else {
		for _, val := range m {
			if val > 1 {
				ans++
			}
		}
	}
	return ans
}
