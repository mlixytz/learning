package leetcode057

func insert(intervals [][]int, newInterval []int) [][]int {
	var left, right [][]int
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			left = append(left, intervals[i])
		} else if intervals[i][0] > newInterval[1] {
			right = append(right, intervals[i])
		} else {
			if intervals[i][0] < newInterval[0] {
				newInterval[0] = intervals[i][0]
			}
			if intervals[i][1] > newInterval[1] {
				newInterval[1] = intervals[i][1]
			}
		}
	}
	return append(append(left, newInterval), right...)
}
