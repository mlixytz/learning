package leetcode048

// 先对称，后转置

func rotate(matrix [][]int) {
	for i, j := 0, len(matrix)-1; i < j; {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
