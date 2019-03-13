package leetcode268

/*
   相同数字 异或 为0
   0 与任何数字 异或 为数字本身
*/

func missingNumber(nums []int) int {
	var ans int
	for i, val := range nums {
		ans ^= i ^ val
	}
	return ans ^ len(nums)
}
