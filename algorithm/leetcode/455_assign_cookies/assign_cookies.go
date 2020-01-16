package leetcode455

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var i, j, res int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
			res++
		} else {
			j++
		}
	}
	return res
}
