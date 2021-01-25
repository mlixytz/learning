package leetcode769

func maxChunksToSorted(arr []int) int {
	res, curMax := 0, arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > curMax {
			curMax = arr[i]
		}
		if i >= curMax {
			res++
		}
	}
	return res
}
