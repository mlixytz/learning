package leetcode056

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	res := make([]Interval, 0)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })
	start := intervals[0].Start
	end := intervals[0].End
	for _, ins := range intervals {
		if ins.Start <= end {
			if ins.End > end {
				end = ins.End
			}
		} else {
			res = append(res, Interval{start, end})
			start, end = ins.Start, ins.End
		}
	}
	res = append(res, Interval{start, end})
	return res
}
