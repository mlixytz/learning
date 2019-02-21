package leetcode074

func searchMatrix(matrix [][]int, target int) bool {
	matLen := len(matrix)
	if matLen == 0 || len(matrix[0]) == 0 {
		return false
	}
	l, h := 0, matLen-1

	for l < h {
		m := l + (h-l)/2
		indexVal := len(matrix[m]) - 1
		if matrix[m][indexVal] < target {
			l = m + 1
		} else if matrix[m][indexVal] > target {
			h = m
		} else {
			return true
		}
	}
	if matrix[h][len(matrix[h])-1] < target || matrix[h][0] > target {
		return false
	}
	searchIndex := h
	l, h = 0, len(matrix[searchIndex])-1
	for l <= h {
		m := l + (h-l)/2
		if matrix[searchIndex][m] < target {
			l = m + 1
		} else if matrix[searchIndex][m] > target {
			h = m - 1
		} else {
			return true
		}
	}
	return false
}
