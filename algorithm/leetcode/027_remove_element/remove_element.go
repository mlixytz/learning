package leetcode027

/*
   使用双指针，一个表示和val不相同的个数。一个遍历的索引
*/

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
