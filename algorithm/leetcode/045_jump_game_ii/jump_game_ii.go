package leetcode045

/* 思路：贪心算法， 每次找这一跳之后可以去的最远距离 */

func jump(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	count, cur := 0, 0
	for cur < n-1 {
		count++
		max, _cur := 0, 0
		for i := 1; i <= nums[cur]; i++ {
			if cur+i == n-1 { // 如果cur+i等于n-1 说明这一跳可以到达列表尾部，所以直接返回。
				return count
			}
			if (cur+i)+nums[cur+i] > max {
				max, _cur = (cur+i)+nums[cur+i], cur+i
			}
		}
		cur = _cur
	}
	return count
}
