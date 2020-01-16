package leetcode001

// my solution
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// best solution
func twoSumBestSolution(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, v := range nums {
		if key, ok := numMap[target-v]; ok {
			return []int{key, i}
		}
		numMap[v] = i
	}
	return []int{1}
}
