package leetcode051

// 普通回溯实现
// func solveNQueens(n int) [][]string {
// 	ans := make([][]string, 0)
// 	tmp := make([][]bool, n)
// 	for i := 0; i < n; i++ {
// 		tmp[i] = make([]bool, n)
// 	}
// 	backtrack(n, 0, &ans, &[]string{}, tmp)
// 	return ans
// }

// func backtrack(n, i int, ans *[][]string, tmpStrs *[]string, tmp [][]bool) {
// 	if i == n {
// 		ts := make([]string, len(*tmpStrs))
// 		copy(ts, *tmpStrs)
// 		*ans = append(*ans, ts)
// 		return
// 	}
// 	for j := 0; j < n; j++ {
// 		if !tmp[i][j] {
// 			_tmp := make([][]bool, n)
// 			for k := 0; k < n; k++ {
// 				_tmp[k] = make([]bool, n)
// 				copy(_tmp[k], tmp[k])
// 			}
// 			var s string
// 			for k := 0; k < n; k++ {
// 				_tmp[i][k] = true // row
// 				_tmp[k][j] = true // col
// 				if i+k < n && j+k < n {
// 					_tmp[i+k][j+k] = true // right down
// 				}
// 				if i-k >= 0 && j-k >= 0 {
// 					_tmp[i-k][j-k] = true // left up
// 				}
// 				if i-k >= 0 && j+k < n {
// 					_tmp[i-k][j+k] = true // right up
// 				}
// 				if i+k < n && j-k >= 0 {
// 					_tmp[i+k][j-k] = true // left down
// 				}
// 				if k == j {
// 					s += "Q"
// 				} else {
// 					s += "."
// 				}
// 			}
// 			*tmpStrs = append(*tmpStrs, s)
// 			backtrack(n, i+1, ans, tmpStrs, _tmp)
// 			*tmpStrs = (*tmpStrs)[:len(*tmpStrs)-1]
// 		}
// 	}
// }

// 按照行遍历，使用三个列表，分别记录 列，45度对角线，135度对角线 是否存在皇后
var (
	col    []bool
	dia45  []bool
	dia135 []bool
	ret    [][]string
)

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return ret
	}

	ret = make([][]string, 0)
	if n == 1 {
		ret = append(ret, []string{"Q"})
		return ret
	}

	col, dia45, dia135 = make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	var slove []string
	putQueen(0, slove)

	return ret
}

func putQueen(x int, slove []string) {
	if len(slove) == len(col) {
		ret = append(ret, slove)
		return
	}

	for y := 0; y < len(col); y++ {
		// 该列已经放置过皇后了,应该找下一列元素来放置
		if col[y] || dia45[x+y] || dia135[x-y+len(col)-1] {
			continue
		} else {
			col[y], dia45[x+y], dia135[x-y+len(col)-1] = true, true, true

			// 更新slove，并从下一行开始找一个位置来放置皇后
			newSlove := make([]string, len(slove)+1)
			copy(newSlove, slove)
			colStr := genarateColStr(y, len(col))
			newSlove[len(slove)] = colStr

			// 递归放置下一行
			putQueen(x+1, newSlove)
			col[y], dia45[x+y], dia135[x-y+len(col)-1] = false, false, false
		}
	}

	return
}

// 生成类似"...Q"的字符串
func genarateColStr(pos int, n int) string {
	byteStr := make([]byte, n)
	for i := 0; i < n; i++ {
		byteStr[i] = '.'
	}

	byteStr[pos] = 'Q'

	return string(byteStr)
}
