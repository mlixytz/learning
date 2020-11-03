package leetcode435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	n := len(intervals)
	res := 0
	prev := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] < prev {
			res++
		} else {
			prev = intervals[i][1]
		}
	}
	return res
}
