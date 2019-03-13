package leetcode078

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	tmpList := make([]int, 0)
	backtrack(nums, &tmpList, &res, 0)
	return res
}

func backtrack(nums []int, tmpList *[]int, res *[][]int, i int) {
	t := make([]int, len(*tmpList))
	copy(t, *tmpList)
	*res = append(*res, t)
	for j := i; j < len(nums); j++ {
		*tmpList = append(*tmpList, nums[j])
		backtrack(nums, tmpList, res, j+1)
		*tmpList = (*tmpList)[:len(*tmpList)-1]
	}
}
