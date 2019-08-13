package leetcode080

/* k和val 两个指针 当当前的数值不比他前两个数值大的时候，意味着出现了3个
或3个以上的相同值 */

func removeDuplicates(nums []int) int {
	k := 0
	for _, val := range nums {
		if k < 2 || val > nums[k-2] {
			nums[k] = val
			k++
		}
	}
	return k
}
