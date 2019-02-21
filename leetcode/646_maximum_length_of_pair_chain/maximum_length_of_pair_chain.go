package leetcode646

import (
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	prev := pairs[0]
	count := 1
	for i := 1; i < len(pairs); i++ {
		if prev[1] < pairs[i][0] {
			count++
			prev = pairs[i]
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
