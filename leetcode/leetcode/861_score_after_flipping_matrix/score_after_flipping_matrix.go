package leetcode861

func matrixScore(A [][]int) int {
	N, M := len(A), len(A[0])
	for i := 0; i < N; i++ {
		if A[i][0] != 1 {
			for j := 0; j < M; j++ {
				A[i][j] = 1 - A[i][j]
			}
		}
	}

	for j := 1; j < M; j++ {
		sum := 0
		for i := 0; i < N; i++ {
			if A[i][j] == 0 {
				sum--
			} else {
				sum++
			}
		}
		if sum < 0 {
			for i := 0; i < N; i++ {
				A[i][j] = 1 - A[i][j]
			}
		}
	}

	sum := 0
	for i := 0; i < N; i++ {
		power := 1
		for j := 0; j < M; j++ {
			sum += A[i][M-j-1] * power
			power *= 2
		}
	}
	return sum
}
