package leetcode052

func totalNQueens(n int) int {
	col := make([]bool, n)
	dia45 := make([]bool, 2*n-1)
	dia135 := make([]bool, 2*n-1)
	var count int

	backtrack(0, n, &col, &dia45, &dia135, &count)
	return count
}

func backtrack(row, n int, col, dia45, dia135 *[]bool, count *int) {
	if row == n {
		(*count)++
		return
	}
	for i := 0; i < n; i++ {
		if (*col)[i] || (*dia45)[row+i] || (*dia135)[n+row-i-1] {
			continue
		}
		(*col)[i] = true
		(*dia45)[row+i] = true
		(*dia135)[n+row-i-1] = true
		backtrack(row+1, n, col, dia45, dia135, count)
		(*col)[i] = false
		(*dia45)[row+i] = false
		(*dia135)[n+row-i-1] = false
	}
}
