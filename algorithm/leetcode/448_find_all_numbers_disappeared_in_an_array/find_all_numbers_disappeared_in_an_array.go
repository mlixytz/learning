package leetcode448

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val < 0 {
			val = -val
		}
		index := val - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	res := make([]int, 0)
	for j, v := range nums {
		if v > 0 {
			res = append(res, j+1)
		}
	}
	return res
}
