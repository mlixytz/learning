package leetcode200

func numIslands(grid [][]byte) int {
	res, m, n := 0, len(grid), len(grid[0])
	var findAllLands func(int, int)
	findAllLands = func(i, j int) {
		grid[i][j] = '0'
		if i+1 < m && grid[i+1][j] == '1' {
			findAllLands(i+1, j)
		}
		if j+1 < n && grid[i][j+1] == '1' {
			findAllLands(i, j+1)
		}
		if i-1 >= 0 && grid[i-1][j] == '1' {
			findAllLands(i-1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			findAllLands(i, j-1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				findAllLands(i, j)
			}
		}
	}

	return res
}
