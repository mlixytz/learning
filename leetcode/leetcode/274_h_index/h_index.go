package leetcode274

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}
	return 0
}
