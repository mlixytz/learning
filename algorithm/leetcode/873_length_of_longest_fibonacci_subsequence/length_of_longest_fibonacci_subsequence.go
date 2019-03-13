package leetcode873

func lenLongestFibSubseq(A []int) int {
	_max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	size := len(A)
	ml := 0
	m := make(map[int]int, size)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		m[A[i]] = i
		for j := 0; j < i; j++ {
			if k, ok := m[A[i]-A[j]]; ok && k < j {
				dp[j][i] = dp[k][j] + 1
			} else {
				dp[j][i] = 2
			}
			ml = _max(dp[j][i], ml)
		}
	}
	if ml > 2 {
		return ml
	}
	return 0
}

// best solution
func lenLongestFibSubseqBS(A []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	size := len(A)
	res := 0
	for k := 2; k < size; k++ {
		for i, j := 0, k-1; i < j; {
			s := A[i] + A[j]

			if s < A[k] {
				i++
				continue
			} else if s > A[k] {
				j--
				continue
			}

			count, x0, x1, next := 3, A[j], A[k], k+1
			for {
				l, h, x2 := next, size-1, x0+x1
				for l <= h {
					m := l + (h-l)/2
					if A[m] > x2 {
						h = m - 1
					} else if A[m] < x2 {
						l = m + 1
					} else {
						count, x0, x1, next = count+1, x1, x2, m+1
						break
					}
				}
				if l > h || next > size-1 {
					break
				}
			}
			res = max(res, count)
			i++
			j--
		}
	}
	return res
}
