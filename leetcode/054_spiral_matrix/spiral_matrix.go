package leetcode054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	rightBorder, downBorder, leftBorder, topBorder := len(matrix[0])-1, len(matrix)-1, 0, 0
	var direct, x, y int
	for topBorder <= downBorder && leftBorder <= rightBorder {
		if direct == 0 { // 向右
			res = append(res, matrix[x][y])
			if y == rightBorder {
				x++
				direct = 1
				topBorder++
			} else {
				y++
			}
		} else if direct == 1 { // 向下
			res = append(res, matrix[x][y])
			if x == downBorder {
				y--
				direct = 2
				rightBorder--
			} else {
				x++
			}
		} else if direct == 2 { // 向左
			res = append(res, matrix[x][y])
			if y == leftBorder {
				x--
				direct = 3
				downBorder--
			} else {
				y--
			}
		} else if direct == 3 { // 向上
			res = append(res, matrix[x][y])
			if x == topBorder {
				y++
				direct = 0
				leftBorder++
			} else {
				x--
			}
		}
	}
	return res
}
