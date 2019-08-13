package leetcode874

func robotSim(commands []int, obstacles [][]int) int {
	var i, j, res int
	xy := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	obSet := make(map[int]bool)
	for _, val := range obstacles {
		obSet[makeKey(val[0], val[1])] = true
	}
	direct := 0
	for _, c := range commands {
		switch c {
		case -1:
			direct++
		case -2:
			direct--
			if direct < 0 {
				direct = 4 + direct
			}
		default:
			d := direct % 4
			dx, dy := xy[d][0], xy[d][1]
			for c > 0 {
				if _, ok := obSet[makeKey(i+dx, j+dy)]; ok {
					break
				}
				c--
				i += dx
				j += dy
			}
			res = max(res, i*i+j*j)
		}
	}
	return res
}

func makeKey(a int, b int) int {
	return a*30000 + b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
