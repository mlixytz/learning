package leetcode169

func majorityElement(nums []int) int {
	majElem, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			majElem = nums[i]
		} else if nums[i] == majElem {
			count++
		} else {
			count--
		}
	}
	return majElem
}
