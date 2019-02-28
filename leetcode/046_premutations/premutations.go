package leetcode046

func permute(nums []int) [][]int {
	list := make([][]int, 0)
	backtrack(&list, 0, &nums)
	return list
}

func backtrack(list *[][]int, pos int, nums *[]int) {
	if pos == len(*nums) {
		t := make([]int, len(*nums))
		copy(t, *nums)
		*list = append(*list, t)
		return
	}
	for i := pos; i < len(*nums); i++ {
		(*nums)[i], (*nums)[pos] = (*nums)[pos], (*nums)[i]
		backtrack(list, pos+1, nums)
		(*nums)[i], (*nums)[pos] = (*nums)[pos], (*nums)[i]
	}
}
