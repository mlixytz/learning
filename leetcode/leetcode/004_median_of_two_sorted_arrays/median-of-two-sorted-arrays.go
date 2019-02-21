package leetcode004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len, i := len(nums1), 0
	nums2Len, j := len(nums2), 0

	resNums := make([]int, nums1Len+nums2Len)

	for k := 0; k < nums1Len+nums2Len; k++ {
		if j == nums2Len || (i < nums1Len && j < nums2Len && nums1[i] <= nums2[j]) {
			resNums[k] = nums1[i]
			i++
			continue
		}
		if i == nums1Len || (i < nums1Len && j < nums2Len && nums1[i] > nums2[j]) {
			resNums[k] = nums2[j]
			j++
		}
	}

	resLen := len(resNums)
	if resLen%2 == 0 {
		return float64(resNums[resLen/2-1]+resNums[resLen/2]) / 2.0
	}
	return float64(resNums[resLen/2])
}
