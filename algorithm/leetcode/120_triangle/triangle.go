package leetcode120

import (
	"math"
)

// my solution
func minimumTotal(triangle [][]int) int {
	size := len(triangle)
	sum := make([]int, size)
	sum[0] = 0
	minVal := math.MaxInt32
	for i := 0; i < size; i++ {
		temp := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				temp[j] = sum[j] + triangle[i][j]
			} else if j == i {
				temp[j] = sum[j-1] + triangle[i][j]
			} else {
				temp[j] = min(sum[j-1], sum[j]) + triangle[i][j]
			}
			if i == size-1 {
				minVal = min(minVal, temp[j])
			}
		}
		sum = temp
	}
	return minVal
}

// best solution
func minimumTotalBS(triangle [][]int) int {
	size := len(triangle)
	sum := triangle[size-1]
	for i := size - 2; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			sum[j] = min(sum[j], sum[j+1]) + triangle[i][j]
		}
	}
	return sum[0]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
