package leetcode1051

import "sort"

func heightChecker(heights []int) int {
	if len(heights) < 2 {
		return 0
	}
	newHeights := make([]int, len(heights))
	copy(newHeights, heights)
	sort.Ints(heights)
	ret := 0
	for i, val := range newHeights {
		if val != heights[i] {
			ret++
		}
	}
	return ret
}
