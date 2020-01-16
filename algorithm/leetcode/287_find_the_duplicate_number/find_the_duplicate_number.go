package leetcode287

/* 与Linked List Cycle II（142）相同，将对应索引的值看成是下个元素的索引
   有重复数据一定会产生圆环
*/

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}
