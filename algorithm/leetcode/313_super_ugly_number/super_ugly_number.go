package leetcode313

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, n)
	index := make([]int, len(primes)) // 代表每个prime最小可以乘的ugly的索引（去重）
	ugly[0] = 1
	for i := 1; i < n; i++ {
		// 找到下一个
		ugly[i] = math.MaxInt64
		for j := 0; j < len(primes); j++ {
			if primes[j]*ugly[index[j]] == ugly[i-1] {
				index[j]++
			}
			ugly[i] = min(ugly[i], ugly[index[j]]*primes[j])
		}
	}
	return ugly[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
