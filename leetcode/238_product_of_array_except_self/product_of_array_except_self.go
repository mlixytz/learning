package leetcode238

/* 执行两轮
   第一轮：从左至右，使所有当前元素对应的结果都是其左边所有元素的乘积
   第二轮：从右至左，使所有当前元素对应的结果都是其右边所有元素的乘积 * 当前元素第一轮结果
*/

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}
	return ans
}
