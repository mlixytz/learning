package leetcode036

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row, column, cube := make(map[byte]bool), make(map[byte]bool), make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}

			if board[j][i] != '.' {
				if column[board[j][i]] {
					return false
				}
				column[board[j][i]] = true
			}

			// 转换索引，匹配cube
			rowIndex, colIndex := 3*(i/3), 3*(i%3)
			if board[rowIndex+j/3][colIndex+j%3] != '.' {
				if cube[board[rowIndex+j/3][colIndex+j%3]] {
					return false
				}
				cube[board[rowIndex+j/3][colIndex+j%3]] = true
			}
		}
	}
	return true
}
