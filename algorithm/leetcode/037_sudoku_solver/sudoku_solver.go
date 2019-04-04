package leetcode037

func solveSudoku(board [][]byte) {
	tmpboard := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		tmpboard[i] = make([]byte, 9)
		copy(tmpboard[i], board[i])
	}
	backtrack(tmpboard, 0, board)
}

func backtrack(tmpboard [][]byte, i int, board [][]byte) {
	if i == 81 {
		for i := 0; i < 9; i++ {
			copy(board[i], tmpboard[i])
		}
		return
	}
	row, col := i/9, i%9
	if tmpboard[row][col] == '.' {
		restNums := findRestNums(row, col, tmpboard)
		if len(restNums) == 0 {
			return
		}
		for _, val := range restNums {
			tmpboard[row][col] = val
			backtrack(tmpboard, i+1, board)
		}
		tmpboard[row][col] = '.'
	} else {
		backtrack(tmpboard, i+1, board)
	}
}

// 找到可以使用的数字
func findRestNums(row, col int, board [][]byte) []byte {
	check := [10]bool{}
	for i := 0; i < 9; i++ {
		// row
		if board[i][col] != '.' {
			check[board[i][col]-'0'] = true
		}
		// col
		if board[row][i] != '.' {
			check[board[row][i]-'0'] = true
		}
		// cube
		_row, _col := 3*(row/3)+(i/3), 3*(col/3)+(i%3)
		if board[_row][_col] != '.' {
			check[board[_row][_col]-'0'] = true
		}
	}
	ret := []byte{}
	for i := 1; i < 10; i++ {
		if !check[i] {
			ret = append(ret, byte(i+48))
		}
	}
	return ret
}
