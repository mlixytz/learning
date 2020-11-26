package leetcode077

func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	var backtrack func(int, []int)
	backtrack = func(cur int, tmpList []int) {
		if len(tmpList) == k {
			t := make([]int, len(tmpList))
			copy(t, tmpList)
			res = append(res, t)
			return
		}
		for i := cur; i <= n; i++ {
			tmpList = append(tmpList, i)
			backtrack(i+1, tmpList)
			tmpList = tmpList[:len(tmpList)-1]
		}
	}

	return res
}
