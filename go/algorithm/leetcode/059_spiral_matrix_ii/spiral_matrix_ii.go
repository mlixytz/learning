package leetcode059

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	num, rightBorder, downBorder, leftBorder, topBorder := 1, n-1, n-1, 0, 0
	var direct, x, y int
	for topBorder <= downBorder && leftBorder <= rightBorder {
		if direct == 0 { // 向右
			matrix[x][y] = num
			num++
			if y == rightBorder {
				x++
				direct = 1
				topBorder++
			} else {
				y++
			}
		} else if direct == 1 { // 向下
			matrix[x][y] = num
			num++
			if x == downBorder {
				y--
				direct = 2
				rightBorder--
			} else {
				x++
			}
		} else if direct == 2 { // 向左
			matrix[x][y] = num
			num++
			if y == leftBorder {
				x--
				direct = 3
				downBorder--
			} else {
				y--
			}
		} else if direct == 3 { // 向上
			matrix[x][y] = num
			num++
			if x == topBorder {
				y++
				direct = 0
				leftBorder++
			} else {
				x--
			}
		}
	}
	return matrix
}
