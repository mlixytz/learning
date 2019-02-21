package leetcode350

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, val := range nums1 {
		m[val]++
	}

	for _, val := range nums2 {
		if count, ok := m[val]; ok && count > 0 {
			m[val]--
			res = append(res, val)
		}
	}
	return res
}
