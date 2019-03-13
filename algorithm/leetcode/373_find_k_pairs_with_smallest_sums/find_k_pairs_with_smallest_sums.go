package leetcode373

/*
   ?????????????????????????
   ???????????????????????????
*/

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1)*len(nums2) < k {
		k = len(nums1) * len(nums2)
	}
	ans := make([][]int, 0)
	t := make([]int, len(nums1)) // t????
	for k > 0 {
		min := 2 << 31
		minIndex := 0
		for i := 0; i < len(nums1); i++ {
			if t[i] > len(nums2)-1 {
				continue
			}
			if nums1[i]+nums2[t[i]] < min {
				min = nums1[i] + nums2[t[i]]
				minIndex = i
			}
		}
		ans = append(ans, []int{nums1[minIndex], nums2[t[minIndex]]})
		t[minIndex]++
		k--
	}
	return ans
}
