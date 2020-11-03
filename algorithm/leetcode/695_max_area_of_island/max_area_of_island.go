package leetcode695

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			tmp := maxArea(grid, i, j)
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

func maxArea(grid [][]int, i int, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
		grid[i][j] = 0
		return 1 + maxArea(grid, i+1, j) + maxArea(grid, i-1, j) + maxArea(grid, i, j-1) + maxArea(grid, i, j+1)
	}
	return 0
}
