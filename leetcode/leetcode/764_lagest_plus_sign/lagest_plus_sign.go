package leetcode764

func orderOfLargestPlusSign(N int, mines [][]int) int {
	maxtri := [500][500]bool{}
	for _, v := range mines {
		maxtri[v[0]][v[1]] = true
	}
	order := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if !maxtri[i][j] {
				k := 1
				for true {
					if i+k > N-1 || i-k < 0 || j+k > N-1 || j-k < 0 {
						break
					}

					// left
					if maxtri[i][j-k] {
						break
					}

					// right
					if maxtri[i][j+k] {
						break
					}

					// up
					if maxtri[i-k][j] {
						break
					}

					// down
					if maxtri[i+k][j] {
						break
					}
					k++
				}
				order = max(k, order)
			}
		}
	}
	return order
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
