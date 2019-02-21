package leetcode475

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	res := 0
	for _, house := range houses {
		l, h, index := 0, len(heaters)-1, -1
		for l <= h {
			m := l + (h-l)/2
			if house > heaters[m] {
				l = m + 1
			} else if house < heaters[m] {
				h = m - 1
			} else {
				index = m
				break
			}
		}
		if index >= 0 {
			continue
		}

		distLeft, distRight := 0, 0
		if l-1 < 0 {
			distLeft = house - heaters[0]
		} else {
			distLeft = house - heaters[l-1]
		}

		if l >= len(heaters) {
			distRight = house - heaters[l-1]
		} else {
			distRight = heaters[l] - house
		}

		res = max(res, min(abs(distLeft), abs(distRight)))
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
