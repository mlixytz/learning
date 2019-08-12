package leetcode060

import "strconv"

func getPermutation(n int, k int) string {
	nums := make(map[int]int)
	nums[0] = 1
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] * i
	}

	k--
	var ret string
	deleted := make(map[int]bool)

	for i := n; i > 0; i-- {
		for j, l := 0, 0; j < n; j++ {
			if _, ok := deleted[j]; ok {
				continue
			}
			if l == k/nums[i-1] {
				deleted[j] = true
				ret += strconv.Itoa(j + 1)
				break
			}
			l++
		}
		k %= nums[i-1]
	}
	return ret
}
