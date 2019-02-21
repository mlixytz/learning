package leetcode062

// my solution
func uniquePaths(m int, n int) int {
	ways := make([][]int, n)
	for index := range ways {
		ways[index] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j > 0 {
				ways[i][j] = ways[i][j-1]
			} else if j == 0 && i > 0 {
				ways[i][j] = ways[i-1][j]
			} else if i == 0 && j == 0 {
				ways[i][j] = 1
			} else {
				ways[i][j] = ways[i-1][j] + ways[i][j-1]
			}
		}
	}
	return ways[n-1][m-1]
}

// best solution
func uniquePathsBS(m int, n int) int {
	ways := make([]int, n)
	for index := range ways {
		ways[index] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ways[j] += ways[j-1]
		}
	}
	return ways[n-1]
}
