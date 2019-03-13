package leetcode349

func intersection(nums1 []int, nums2 []int) []int {
	h1 := make(map[int]bool)
	h2 := make(map[int]bool)
	res := make([]int, 0)
	for _, val := range nums1 {
		if _, ok := h1[val]; !ok {
			h1[val] = true
		}
	}

	for _, val := range nums2 {
		if _, ok1 := h1[val]; ok1 {
			if _, ok2 := h2[val]; !ok2 {
				res = append(res, val)
				h2[val] = true
			}
		}
	}
	return res
}
